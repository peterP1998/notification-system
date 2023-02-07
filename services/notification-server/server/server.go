package server

func Init(host string, kafkaHost string) {
	r := InitRouter(kafkaHost)
	r.Run(host)
}
