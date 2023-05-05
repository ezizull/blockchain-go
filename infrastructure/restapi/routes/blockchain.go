package routes

import (
	blockController "blockchain-go/infrastructure/restapi/controllers/blockchain"

	"github.com/gofiber/fiber/v2"
)

// BlockChainRoutes is a function that contains all routes of the block chain
func BlockChainRoutes(router fiber.Router, controller *blockController.Controller) {
	routerBlockChain := router.Group("/block-chain")

	// authentication
	// routerPhoto.Use(middlewares.AuthJWTMiddleware())
	{
		routerBlockChain.Get("", controller.GetAllBlockByPort)
	}

}
