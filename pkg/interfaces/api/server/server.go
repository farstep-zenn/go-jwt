package server

import (
	"log"
	"net/http"

	"github.com/FarStep131/go-jwt/pkg/infrastructure"
	"github.com/FarStep131/go-jwt/pkg/infrastructure/repositoryimpl"
	"github.com/FarStep131/go-jwt/pkg/interfaces/api/handler"
	"github.com/FarStep131/go-jwt/pkg/interfaces/api/middleware"
	"github.com/FarStep131/go-jwt/pkg/usecase"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// サーバ起動処理
func Serve(addr string) {
	// 依存性の注入
	userRepoImpl := repositoryimpl.NewRepositoryImpl(infrastructure.Conn)
	userUseCase := usecase.NewUseCase(userRepoImpl)
	userHandler := handler.NewHandler(userUseCase)

	r = gin.Default()

	r.POST("/signup", userHandler.HandleSignup)
	r.POST("/login", userHandler.HandleLogin)
	r.GET("/logout", userHandler.HandleLogout)

	secured := r.Group("/secured").Use(middleware.Auth())
	secured.GET("/ping", Ping)

	log.Println("Server running...")
	if err := r.Run(addr); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
