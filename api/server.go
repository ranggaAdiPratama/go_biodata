package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/ranggaAdiPratama/go_biodata/db/sqlc"
	"github.com/ranggaAdiPratama/go_biodata/token"
	"github.com/ranggaAdiPratama/go_biodata/util"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)

	if err != nil {
		return nil, fmt.Errorf("Cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// authRoutes.GET("/accounts", server.listAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.POST("/profile", server.updateProfile)
	// authRoutes.POST("/transfers", server.createTransfer)

	router.POST("/auth/register", server.register)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
