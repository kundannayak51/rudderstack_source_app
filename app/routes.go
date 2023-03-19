package app

import (
	"github.com/rudderstack_source_app/app/controller"
	"github.com/rudderstack_source_app/repository/source"
	"github.com/rudderstack_source_app/repository/user"
	sourceservice "github.com/rudderstack_source_app/services/source"
	"github.com/rudderstack_source_app/services/sourcetemplate"
	userservice "github.com/rudderstack_source_app/services/user"

	"github.com/rudderstack_source_app/repository/source_template"

	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(client *mongo.Client) {
	sourceRepo := source.NewSourceRepository(client)
	sourceTemplateRepo := source_template.NewSourceTemplateRepository(client)
	userRepo := user.NewUserRepository(client)

	sourceTemplateService := sourcetemplate.NewService(sourceTemplateRepo, userRepo)
	sourceService := sourceservice.NewService(sourceRepo, sourceTemplateRepo, userRepo)
	userService := userservice.NewService(userRepo)

	sourceTemplateController := controller.NewSourceTemplateController(sourceTemplateService)
	sourceController := controller.NewSourceController(sourceService)
	userController := controller.NewUserController(userService)

	Router.POST("/source-template", sourceTemplateController.AddSourceTemplate)
	Router.GET("/source-template/:type", sourceTemplateController.GetSourceTemplateByType)
	Router.GET("/all-source-types", sourceTemplateController.GetAllSourceTypes)

	Router.POST("/source", sourceController.CreateSource)
	Router.GET("/sources/:type", sourceController.GetAllSources)

	Router.POST("/user", userController.CreateOrUpdateUser)
}
