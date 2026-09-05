package routers

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/ye-f-ying/dark_back/app/middlewares"
)

func Router(h *server.Hertz) {
	if h == nil {
		return
	}
	h.Use(middlewares.CorsMiddleware())
	adminRouter := h.Group("/admin")
	adminRouters(adminRouter)
}
