package http

import (
	"log"
	"shortURL/internal/delivery/handler"
	"shortURL/pkg/repository/redis"
	"shortURL/pkg/service"

	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

type HttpServer struct {
	*gin.Engine
}

func (server *HttpServer) setRouter(handler *handler.Handler) {
	v1 := server.Group("api/v1")
	{
		v1.POST("/urls", handler.UploadURL)
		v1.DELETE("/urls/:url", handler.DeleteURL)
	}

	server.GET("/:url", handler.RedirectURL)

}

func NewHttpServer(service *service.ShortURL, redis *redis.Cache) *HttpServer {
	httpServer := &HttpServer{
		Engine: gin.Default(),
	}

	handler := &handler.Handler{
		ShortURL: service,
	}

	rate, err := limiter.NewRateFromFormatted("60-H")
	if err != nil {
		log.Fatal(err)
	}

	store, err := sredis.NewStoreWithOptions(redis.Client, limiter.StoreOptions{
		Prefix:   "limiter_gin_example",
		MaxRetry: 3,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Create a new middleware with the limiter instance.
	middleware := mgin.NewMiddleware(limiter.New(store, rate))

	httpServer.Use(middleware)

	httpServer.setRouter(handler)

	return httpServer
}
