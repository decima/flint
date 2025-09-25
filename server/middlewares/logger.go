package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// loggerKey est la clé utilisée pour stocker le logger dans le gin.Context.
// Elle n'est pas exportée pour empêcher les autres packages d'y accéder directement.
const loggerKey = "logger"

type ContextualLoggerMiddleware struct {
	baseLogger *zerolog.Logger
}

func (m *ContextualLoggerMiddleware) Do(c *gin.Context) {
	// Crée un logger enfant avec des champs spécifiques à la requête.
	contextualLogger := m.baseLogger.With().
		Str("ClientIP", c.ClientIP()).
		Str("user-agent", c.Request.UserAgent()).
		Str("referer", c.Request.Referer()).
		Str("host", c.Request.Host).
		Str("path", c.Request.URL.Path).
		Str("method", c.Request.Method).
		Logger()

	// Stocke le nouveau logger dans le contexte pour cette requête.
	c.Set(loggerKey, &contextualLogger)

	// Passe au middleware ou handler suivant.
	c.Next()
}

func NewContextualLoggerMiddleware(baseLogger *zerolog.Logger) *ContextualLoggerMiddleware {
	return &ContextualLoggerMiddleware{baseLogger: baseLogger}
}

// GetLogger récupère le zap.Logger depuis le gin.Context.
// Il paniquera si le logger n'est pas trouvé, car cela indique une erreur de configuration
// (le middleware n'a pas été appliqué).
func GetLogger(c *gin.Context) *zerolog.Logger {
	logger, exists := c.Get(loggerKey)
	if !exists {
		panic("Logger not found in context. Is the middleware correctly applied?")
	}
	// L'assertion de type est sûre car c'est nous qui avons défini la valeur.
	return logger.(*zerolog.Logger)
}

type AccessLogger struct{}

func (a *AccessLogger) Do(c *gin.Context) {
	start := time.Now()
	c.Next()
	duration := time.Since(start)
	statusCode := c.Writer.Status()

	logger := GetLogger(c).With().Dur("duration", duration).Int("status", statusCode).Logger()
	logger.Info().Msg(c.Request.Method + " " + c.Request.RequestURI)
}

func NewAccessLogger() *AccessLogger {
	return &AccessLogger{}
}
