package server

func Init(serverHost string) {
	r := InitRouter()
	r.Run(serverHost)
}
