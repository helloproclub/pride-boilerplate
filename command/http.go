package command

import (
	"fmt"

	"github.com/helloproclub/pride-boilerplate/helper"
	"github.com/helloproclub/pride-boilerplate/http"
	"github.com/helloproclub/pride-boilerplate/repository"
	"github.com/helloproclub/pride-boilerplate/service"
	"github.com/helloproclub/pride-boilerplate/storage"
	"github.com/spf13/cobra"
)

var httpCommand = &cobra.Command{
	Use:   "http",
	Short: "serve as http server",
	Long:  "serve as http server",
	Run:   httpServe,
}

func init() {
	RootCmd.AddCommand(httpCommand)
}

func httpServe(cmd *cobra.Command, args []string) {
	dbHost := fmt.Sprintf("v", helper.GetEnv("DB_HOST", ""))
	dbPort := fmt.Sprintf("v", helper.GetEnv("DB_HOST", ""))
	dbUser := fmt.Sprintf("v", helper.GetEnv("DB_HOST", ""))
	dbName := fmt.Sprintf("v", helper.GetEnv("DB_HOST", ""))
	dbPass := fmt.Sprintf("v", helper.GetEnv("DB_HOST", ""))
	db := storage.InitDB(dbHost, dbPort, dbUser, dbName, dbPass)

	// init repositories
	todoRepo := repository.NewPostgresTodoRepository(db)

	// init service
	service := service.Service{
		Todo: service.NewTodoService(todoRepo),
	}

	// init server
	server := http.Init(service)
	server.Serve()
}
