package web

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Server interface contract...
type Server interface {
	Listen()
	// Shutdown(context.Context)
}

type HTTPServer struct {
	port   string
	router *fiber.App
}

func NewHTTPServer(port string, router *fiber.App) HTTPServer {
	return HTTPServer{
		port:   port,
		router: router,
	}
}

func (h HTTPServer) Listen() {
	log.Print(fmt.Sprintf("Web server started on localhost:%s", h.port))

	// h.router.Listen(":3300")
	h.router.Listen(fmt.Sprintf(":%s", h.port))
}
