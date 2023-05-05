// Package routes contains all routes of the application
package routes

import (
	_ "blockchain-go/docs"
	databsDomain "blockchain-go/domain/database"
	"blockchain-go/infrastructure/restapi/adapter"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func ApplicationRootRouter(router fiber.Router, db databsDomain.Database) {
	// Monitoring
	{
		router.Get("/metrics", monitor.New())
	}

	// Documentation Swagger
	{
		router.Get("/swagger/*any", fiberSwagger.WrapHandler)
		router.Get("/", func(c *fiber.Ctx) error {
			return c.Redirect("/swagger/index.html", fiber.StatusMovedPermanently)
		})
	}
}

func ApplicationV1Router(router fiber.Router, db databsDomain.Database) {
	routerV1 := router.Group("/v1")

	{
		// User Routes
		BlockChainRoutes(routerV1, adapter.BlockChainAdapter(db))
	}
}
