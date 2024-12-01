package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Config interface {
	comfig.Logger
	comfig.Listenerer
}

type config struct {
	comfig.Logger
	comfig.Listenerer

	getter kv.Getter
}

func New(getter kv.Getter) Config {
	logger := comfig.NewLogger(getter, comfig.LoggerOpts{})
	return &config{
		getter:     getter,
		Logger:     logger,
		Listenerer: comfig.NewListenerer(getter),
	}
}
