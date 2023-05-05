package blockChain

import (
	blockDomain "blockchain-go/domain/block"
	redisRepo "blockchain-go/infrastructure/repository/redis"
	"encoding/json"
	"strconv"
	"time"

	useCaseBlockChain "blockchain-go/application/usecases/blockchain"
	mssgConst "blockchain-go/utils/constant/message"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// Controller is a struct that contains the user service
type Controller struct {
	InfoRedis  *redisRepo.InfoDatabaseRedis
	BlockChain useCaseBlockChain.Service
}

// GetAllBlockByPort godoc
// @Tags blockchain
// @Summary Get Blockchain By Port Service
// @Description Create new blockchain on the system
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body GetAllBlockByPort true "body data"
// @Success 200 {object} blockDomain.Blockchain
// @Failure 400 {object} controllers.MessageResponse
// @Failure 500 {object} controllers.MessageResponse
// @Router /blockchain [post]
func (c *Controller) GetAllBlockByPort(ctx *fiber.Ctx) error {
	var blockChain *blockDomain.BlockchainResponse

	redisDB := c.InfoRedis.NewRedis(1)
	redisData, redisErr := redisDB.Get(c.InfoRedis.CTX, ctx.Port()).Result()

	if redisErr == redis.Nil {
		port, err := strconv.ParseUint(ctx.Port(), 10, 64)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": mssgConst.StatusInvalidPort})
		}

		blockChain, err = c.BlockChain.GetAllChainByPort(port)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}

		blockChainJSON, err := json.Marshal(blockChain)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}

		status := redisDB.Set(c.InfoRedis.CTX, ctx.Port(), blockChainJSON, 30*60*time.Second)
		if status.Err() != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed set redis"})
		}

	} else {
		err := json.Unmarshal([]byte(redisData), &blockChain)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}

	}

	return ctx.Status(fiber.StatusOK).JSON(blockChain)
}
