package router

import (
	"neuro-most/auth-service/internal/adapters/api/action"
	"neuro-most/auth-service/internal/usecase"
	"neuro-most/auth-service/internal/utils"

	"github.com/gin-gonic/gin"
)

type RouterHTTP struct {
	jwt    utils.JWKSHandler
	router *gin.Engine
}

func NewRouterHTTP(jwt utils.JWKSHandler) RouterHTTP {
	router := gin.Default()
	return RouterHTTP{
		jwt:    jwt,
		router: router,
	}
}

func (r *RouterHTTP) Listen() {
	r.SetupRoutes()
	r.router.Run(":8080")
}

func (r *RouterHTTP) SetupRoutes() {
	r.router.GET("/jwts", r.buildValidateJwts())
}

func (r *RouterHTTP) buildValidateJwts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc  = usecase.NewFindJwtsInteractor(r.jwt)
			act = action.NewFindJwtsAction(uc)
		)

		act.Execute(c.Writer, c.Request)
	}
}
