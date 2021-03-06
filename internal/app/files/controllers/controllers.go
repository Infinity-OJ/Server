package controllers

import (
	"github.com/infinity-oj/server/internal/pkg/transports/http"
	"github.com/infinity-oj/server/internal/pkg/transports/http/middlewares/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func CreateInitControllersFn(pc *FilesController) http.InitControllers {
	return func(r *gin.Engine) {
		r.POST("/session", pc.CreateSession)
		r.GET("/session", jwt.JWT(), pc.GetSession)
	}
}

var ProviderSet = wire.NewSet(NewFilesController, CreateInitControllersFn)
