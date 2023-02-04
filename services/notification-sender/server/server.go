package server

func Init() {
	r := InitRouter()
	r.Run("localhost:8021")
}