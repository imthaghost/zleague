package proxy

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
	"zleague/api/config"
)

/*

If several goroutines are sending requests,
new connections will be created the pool has all
connections busy and will create new ones.
We limit the maximum number of connections per host.

*/
var (
	MaxIdleConnections = 10         // Max Idle Connections
	once               sync.Once    // sync so we only setup 1 client
	netClient          *http.Client // client
)

/*

Transport has its own connection pool,
by default each connection in that pool is
reused if body is fully read and closed
*/

// newProxy will return a proxy function for use
// in a transport that always returns the same URL
func newProxy() func(*http.Request) (*url.URL, error) {
	// get proxy config
	c := config.GetProxyConfig()

	base := "http://%s:%s@%s"
	// fill credentials into url
	proxyURL := fmt.Sprintf(base, c.Username, c.Password, c.Address)
	// parse proxy url
	link, err := url.Parse(proxyURL)
	if err != nil {
		fmt.Println(err)
	}
	// setup proxy transport
	proxy := http.ProxyURL(link)

	return proxy
}

/*

Transport is the struct that holds connections for re-use;
Creating new Transport for each request, will create new connections each time.
In this case the solution is to share the one Transport instance between clients.
By default, Transport caches connections for future re-use.

Reference: https://godoc.org/net/http#Transport

*/

// NewNetClient creates a new client
func NewNetClient() *http.Client {
	once.Do(func() {
		// transport configuratin
		var netTransport = &http.Transport{
			Proxy:        newProxy(),         // default - rotating IP addresses
			MaxIdleConns: MaxIdleConnections, // max idle connections
			Dial: (&net.Dialer{ // Dialer
				Timeout: 20 * time.Second, // max dialer timeout
			}).Dial,
			TLSHandshakeTimeout: 20 * time.Second, // transport layer security max timeout
		}
		netClient = &http.Client{
			Timeout:   time.Second * 20, // roundtripper timeout
			Transport: netTransport,     // how our HTTP requests are made
		}
	})

	return netClient
}
