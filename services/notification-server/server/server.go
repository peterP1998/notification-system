package server

func Init(host string) {
	r := InitRouter()
	r.Run(host)
}
