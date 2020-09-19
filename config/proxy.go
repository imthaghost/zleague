package config

import "os"

// Proxy configuration variables to establish connection
type ProxyConfig struct {
	Schema   string // Proxy Schema (http, https, socks5)
	Address  string // Proxy address
	Username string // Proxy username
	Password string // Proxy password
	Dynamic  bool   // TODO: Sticky or Dynamic
}

// GetProxyConfig
func GetProxyConfig() *ProxyConfig {
	return &ProxyConfig{
		Schema:   os.Getenv("PROXY_SCHEMA"),
		Address:  os.Getenv("PROXY_ADDRESS"),
		Username: os.Getenv("PROXY_USERNAME"),
		Password: os.Getenv("PROXY_PASSWORD"),
	}
}
