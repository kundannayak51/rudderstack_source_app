package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rudderstack_source_app/database"
)

var (
	Router = gin.Default()
)

func StartApplication(ctx context.Context) {
	client, err := database.ConnectDB(ctx)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer client.Disconnect(ctx)

	SetupRoutes(client)
	Router.Run(":8080")
}
