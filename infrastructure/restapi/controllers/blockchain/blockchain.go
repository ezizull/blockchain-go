package blockChain

import (
	redisRepo "blockchain-go/infrastructure/repository/redis"

	useCaseBlockChain "blockchain-go/application/usecases/blockchain"

	"github.com/gofiber/fiber/v2"
)

// Controller is a struct that contains the user service
type Controller struct {
	InfoRedis  *redisRepo.InfoDatabaseRedis
	BlockChain useCaseBlockChain.Service
}

// GetBlockChainByPort godoc
// @Tags block
// @Summary Create New BlockName
// @Description Create new block on the system
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body GetBlockChainByPort true "body data"
// @Success 200 {object} blockDomain.Block
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /block [post]
func (c *Controller) GetBlockChainByPort(ctx *fiber.Ctx) (err error) {
	// authData := ctx.Locals(authConst.Authorized).(*secureDomain.Claims)

	return ctx.Status(fiber.StatusCreated).JSON(blockChain)
}
