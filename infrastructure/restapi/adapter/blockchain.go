package adapter

import (
	blockChainService "blockchain-go/application/usecases/blockchain"

	databsDomain "blockchain-go/domain/database"

	blockController "blockchain-go/infrastructure/restapi/controllers/blockchain"
)

// BlockChainAdapter is a function that returns a block chain controller
func BlockChainAdapter(db databsDomain.Database) *blockController.Controller {
	service := blockChainService.Service{}

	return &blockController.Controller{
		InfoRedis:  db.Redis,
		BlockChain: service,
	}
}
