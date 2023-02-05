package server

import (
	"github.com/peterP1998/notification-system/notification-server/config"
)

func Init() {
	//config := config.GetConfig()
	r := InitRouter()
	r.Run(config.Configuration.Host)
}
