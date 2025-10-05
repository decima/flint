package main

//func main() {
//
//	rcu := remote.SSHClient{}
//
//	s := model.Server{
//		Host:     "localhost",
//		Port:     22,
//		Username: "decima",
//		Password: "NOPASS",
//		WorkDir:  "~/testflint",
//	}
//	err := rcu.Execute(s, "docker version --format=json", func(stdout io.Reader) error {
//		buf := make([]byte, 1024)
//		for {
//			n, err := stdout.Read(buf)
//			if err != nil {
//				if err == io.EOF {
//					break
//				}
//				return err
//			}
//			fmt.Println("READING...")
//			fmt.Print(string(buf[:n]))
//		}
//
//		return nil
//	})
//	log.Println(err)
//}
