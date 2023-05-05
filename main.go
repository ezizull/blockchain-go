package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	databsDomain "blockchain-go/domain/database"
	"blockchain-go/infrastructure/repository/redis"
	"blockchain-go/infrastructure/restapi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valyala/fasthttp"
)

func main() {
	// initialize config
	router := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	router.Use(logger.New())
	router.Use(limiter.New())
	router.Use(cors.New(cors.ConfigDefault))

	// init databases
	databases := initDatabase()

	// databases routes
	routes.ApplicationV1Router(router, databases)

	// running config
	startServer(router)
}

// initial databases
func initDatabase() databsDomain.Database {
	// redis connection
	redisDB, err := redis.InitRedis()
	if err != nil {
		panic(fmt.Errorf("fatal error in redis: %s", err))
	}

	return databsDomain.Database{
		Redis: redisDB,
	}
}

// start server config
func startServer(app *fiber.App) {
	flagPort := flag.Uint("port", 5000, "custom port number service")
	flag.Parse()

	serverPort := fmt.Sprintf(":%d", *flagPort)
	s := &fasthttp.Server{
		Handler:            app.Handler(),
		ReadTimeout:        18000 * time.Second,
		WriteTimeout:       18000 * time.Second,
		MaxRequestBodySize: 3 << 20,
	}

	if err := s.ListenAndServe(serverPort); err != nil {
		log.Fatalf("fatal error description: %s", strings.ToLower(err.Error()))
	}
}
