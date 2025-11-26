package http

import (
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/infrastructure/security"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/interfaces/http/handler"
	"github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/internal/interfaces/http/middleware"
	"github.com/gin-gonic/gin"

	_ "github.com/cobategit/go-gin-ddd-cleanarchitecture-staterkit/docs" // swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	userHandler *handler.UserHandler,
	jwtService *security.JWTService,
) *gin.Engine {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	users := v1.Group("/users")
	users.Use(middleware.AuthMiddleware(jwtService))
	{
		users.GET("/me", userHandler.Me)
	}

	return r
}
