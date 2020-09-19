package config

import "os"

// ProxyConfig configuration variables to establish connection
type ProxyConfig struct {
	Address  string // Proxy address
	Username string // Proxy username
	Password string // Proxy password
	Dynamic  bool   // TODO: Sticky or Dynamic
}

// GetProxyConfig returns the proxy config
func GetProxyConfig() *ProxyConfig {
	return &ProxyConfig{
		Address:  os.Getenv("PROXY_ADDRESS"),
		Username: os.Getenv("PROXY_USERNAME"),
		Password: os.Getenv("PROXY_PASSWORD"),
	}
}
