package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type ServiceProxy struct {
	serviceName string
	serviceURL  string
}

func NewServiceProxy(serviceName, servicePort string) *ServiceProxy {
	return &ServiceProxy{
		serviceName: serviceName,
		serviceURL:  "http://localhost:" + servicePort,
	}
}

func (s *ServiceProxy) HandleRequest(c *gin.Context) {
	remote, err := url.Parse(s.serviceURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse service URL"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("path")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
