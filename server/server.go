package server

import "quote-me/config"

// Init function initialises the backend
func Init() {
	router := InitRouter()
	router.Run(":" + config.GetConfig().GetString("server.port"))
}
