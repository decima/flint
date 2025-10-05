package handlers

import (
	"flint/security"
	"flint/server/handlers/utils"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
)

type websocketWriter struct {
	conn *websocket.Conn
}

func (w *websocketWriter) Write(p []byte) (n int, err error) {
	err = w.conn.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebsocketRoute struct{}

func NewWebsocketRoute() *WebsocketRoute {
	return &WebsocketRoute{}
}

func (w WebsocketRoute) Route() (utils.Method, utils.Path, *security.Policy) {
	return utils.GET, "/ws", security.AnonymousOrUser()
}

func (w WebsocketRoute) Do(c *gin.Context) {
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer wsConn.Close()
	//  Informations de connexion SSH (à adapter)
	auth := goph.Password("some_password")

	client, err := goph.NewUnknown("decima", "127.0.0.1", auth)
	if err != nil {
		log.Println("Impossible de se connecter au serveur SSH :", err)
		return
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		log.Println("Impossible de créer une session SSH :", err)
		return
	}
	defer session.Close()
	// --- Redirection des flux I/O (la partie importante) ---

	// On récupère les "pipes" pour stdin et stdout.
	// Un "pipe" est un flux de données que l'on peut lire et écrire.
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// On redirige également Stderr vers le même pipe que Stdout.
	session.Stderr = session.Stdout

	go func() {
		// Crée notre writer personnalisé.
		wsWriter := &websocketWriter{conn: wsConn}
		// io.Copy va bloquer et copier les données jusqu'à ce que le pipe stdout se ferme.
		_, err := io.Copy(wsWriter, stdout)
		if err != nil {
			log.Println("Erreur en copiant stdout vers le websocket:", err)
		}
	}()

	go func() {
		defer stdin.Close()
		for {
			// Lit un message du client.
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				log.Println("Erreur de lecture du websocket:", err)
				break // Sort de la boucle si le client se déconnecte.
			}
			// Écrit le message dans le stdin du shell SSH.
			_, err = stdin.Write(message)
			if err != nil {
				log.Println("Erreur d'écriture dans stdin:", err)
				break
			}
		}
	}()

	// --- Configuration du Terminal Virtuel (PTY) ---
	modes := ssh.TerminalModes{
		ssh.ECHO:          1, // Active l'écho des caractères.
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	// Il est crucial de demander un PTY pour avoir un vrai shell interactif.
	if err := session.RequestPty("xterm-256color", 40, 80, modes); err != nil {
		log.Println("La requête PTY a échoué :", err)
		return
	}

	// --- Démarrage du Shell ---
	if err := session.Shell(); err != nil {
		log.Println("Impossible de démarrer le shell :", err)
		return
	}

	// On attend que la session SSH se termine.
	// Cela maintient la fonction handleWebSocket active pendant que les goroutines tournent.
	if err := session.Wait(); err != nil {
		// C'est normal d'avoir une erreur ici quand la session se ferme,
		// donc on peut choisir de l'ignorer ou de la logger sobrement.
		log.Println("Session SSH terminée avec une erreur:", err)
	}
}
