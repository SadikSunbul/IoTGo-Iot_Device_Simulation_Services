package http

import (
	"fmt"
	"github.com/SadikSunbul/IoTGo-Iot_Device_Simulation_Services/pkg/config"
	"github.com/SadikSunbul/IoTGo-Iot_Device_Simulation_Services/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
)

type Server struct {
	app *fiber.App
	cfg *config.Schema
}

func NewServer() *Server {
	return &Server{
		app: fiber.New(),
		cfg: config.GetConfig(),
	}
}

func (s Server) Run() error {
	// Fiber yapılandırması
	s.app = fiber.New(fiber.Config{
		Prefork: s.cfg.Environment == config.ProductionEnv, // Prefork ayarını kontrol et
	})

	// Rotaları eşle
	if err := s.MapRoutes(); err != nil {
		log.Fatalf("MapRoutes Error: %v", err)
	}

	// Swagger dokümantasyonu
	s.app.Get("/swagger/*", swagger.HandlerDefault) // Swagger route

	// Health check endpoint
	s.app.Get("/health", func(c *fiber.Ctx) error {
		response.JSON(c, fiber.StatusOK, "Healthy ❤️❤️❤️")
		return nil
	})

	// HTTP sunucusunu başlat
	log.Printf("HTTP server is listening on PORT: %d", s.cfg.HttpPort)
	if err := s.app.Listen(fmt.Sprintf(":%d", s.cfg.HttpPort)); err != nil {
		log.Fatalf("Running HTTP server: %v", err)
	}

	return nil
}

func (s Server) GetApp() *fiber.App {
	return s.app
}

func (s Server) MapRoutes() error {
	// v1 := s.app.Group("/api/v1")

	return nil
}
