package config

import "github.com/gomodul/envy"

type server struct {
	ADDR string
}

var Server = server{
	envy.Get("SERVER_HOST", "127.0.0.1") + ":" + envy.Get("SERVER_PORT", "80"),
}
