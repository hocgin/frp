package config

import "sync"

type config struct {
	IsDev     bool
	ApiServer string
}

var instance *config
var once sync.Once

func GetConfig() *config {
	once.Do(func() {
		instance = &config{}
	})
	return instance
}

func InitConfig(IsDev bool) {
	GetConfig().IsDev = IsDev

	if instance.IsDev {
		instance.ApiServer = "https://api-dev.hocgin.top/chaos"
	} else {
		instance.ApiServer = "https://api.hocgin.top/chaos"
	}
}
