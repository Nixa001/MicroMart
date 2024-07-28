package handler

import (
	"net/http/httputil"

	"github.com/Nixa001/micromart-gateway/model"
	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	gateway *model.Gateway
}

func NewGatewayHandler(gateway *model.Gateway) *GatewayHandler {
	return &GatewayHandler{gateway: gateway}
}

func (h *GatewayHandler) handleRequest(targetURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy := httputil.NewSingleHostReverseProxy(h.gateway.ParseURL(targetURL))
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func (h *GatewayHandler) HandleProduct(c *gin.Context) {
	h.handleRequest(h.gateway.ProductService)(c)
}

func (h *GatewayHandler) HandleCart(c *gin.Context) {
	h.handleRequest(h.gateway.CartService)(c)
}

func (h *GatewayHandler) HandlePayment(c *gin.Context) {
	h.handleRequest(h.gateway.PaymentService)(c)
}

func (h *GatewayHandler) HandleUser(c *gin.Context) {
	h.handleRequest(h.gateway.UserService)(c)
}
