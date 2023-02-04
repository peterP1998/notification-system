package server 

func Init() {
	//config := config.GetConfig()
	r := InitRouter()
	r.Run("localhost:8020")
}