package server

import "goTimisoaraBackend/config"

func Init() {
	appConfig := config.GetConfig()
	r := NewRouter()

	_ = r.Run(appConfig.GetString("server.port"))
}
