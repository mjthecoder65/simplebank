package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/mjthecoder65/simplebank/db/sqlc"
	"github.com/mjthecoder65/simplebank/token"
	"github.com/mjthecoder65/simplebank/util"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
	maker  token.Maker
	config util.Config
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PasetoSecretKey)

	if err != nil {
		log.Fatal("failed to get token maker", err)
		return nil, err
	}

	server := &Server{
		store:  store,
		maker:  tokenMaker,
		config: config,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter()

	return server, nil
}

func (server *Server) setUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", server.CheckHealth)
	router.GET("/readiness", server.CheckReadiness)

	router.POST("/users", server.createUser)
	router.POST("users/login", server.login)

	authRouters := router.Group("/").Use(AuthMiddleware(server.maker))
	authRouters.POST("/accounts", server.CreateAccount)
	authRouters.GET("/accounts/:id", server.getAccount)
	authRouters.GET("/accounts", server.listAccounts)

	authRouters.POST("/transfers", server.CreateTransfer)
	server.router = router

	return router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
