package server

import (
	"github.com/peterP1998/notification-system/notification-sender/config"
)

func Init() {
	r := InitRouter()
	r.Run(config.Configuration.Host)
}
