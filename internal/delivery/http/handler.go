package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suchimauz/file-storage/internal/config"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	// Init gin handler
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router, cfg)

	return router
}

func (h *Handler) initAPI(router *gin.Engine, cfg *config.Config) {
	group := router.Group("/" + cfg.Storage.Prefix)
	{
		group.GET("/download/:path", func(c *gin.Context) {
			newResponse(c, http.StatusOK, newUploadResponse(c.Param("path")))

			return
		})
		group.POST("/upload/:file", func(c *gin.Context) {})
	}
}
