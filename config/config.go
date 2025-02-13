package config

type RemoteServerConfig struct {
	Username string `json:"username"`
	Host     string `json:"host"`
}

type Config struct {
	RemoteServer []RemoteServerConfig `json:"remoteServer"`
}
