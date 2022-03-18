package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/api/controller"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/api/middleware"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/utils/helper"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/configs"
	"github.com/gofiber/fiber/v2/middleware/cors"
	logs "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"

	_ "github.com/ahmetcanaydemir/go-rest-clean-boilerplate/docs"
)

var rootCmd = &cobra.Command{
	Use:   "go run main.go",
	Short: "Go Rest Clean Boilerplate Service",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.Flags().StringVarP(&configs.Server.Config.Env, "env", "e", "development", "select your environment")
	rootCmd.Flags().StringVarP(&configs.Server.Config.Port, "port", "p", "8080", "api port")

	loc, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		panic(fmt.Sprintf("Error loading location: %s", err))
	}
	configs.Server.Config.Location = loc

	rootCmd.RunE = start
}

// @title                    Go Rest Clean Boilerplate
// @version                  1.0
// @description              Go Rest Clean Boilerplate
// @license.name             Apache 2.0
// @license.url              http://www.apache.org/licenses/LICENSE-2.0.html
// @query.collection.format  pipes
// @produce                  json
// @BasePath                 /
// @schemes                  http
func start(_ *cobra.Command, _ []string) error {

	// region Logger ####
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	zap.ReplaceGlobals(logger)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(fmt.Sprintf("can't sync logger: %v", err))
		}
	}(logger)
	// endregion

	helper.CustomRegistry = helper.CreateCustomRegistry().Build()

	// region Configs ####
	configs.SetConfigs([]byte("{\"DbConnectionString\":\"mongodb://admin:admin@mongodb:27017\"}"))
	// endregion

	// region Fiber ####
	app := fiber.New(fiber.Config{
		GETOnly:      true,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: middleware.ErrorHandler,
	})
	configs.Server.App = app

	// Swagger
	configs.Server.App.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Middlewares
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logs.New(logs.Config{
		Next: middleware.SkipStatusLog,
	}))
	// endregion

	// region Controllers ####
	controller.NewStatusController()
	controller.NewStatsController()
	// endregion

	configs.Server.App.Use(middleware.NotFound)
	err = configs.Server.App.Listen(":" + configs.Server.Config.Port)

	return err
}
