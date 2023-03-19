package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rudderstack_source_app/entity"
	"github.com/rudderstack_source_app/services/sourcetemplate"
	"github.com/rudderstack_source_app/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SourceTemplateController struct {
	sourceTemplateService sourcetemplate.SourceTemplateServiceInterface
}

func NewSourceTemplateController(sourceTemplateService sourcetemplate.SourceTemplateServiceInterface) *SourceTemplateController {
	return &SourceTemplateController{
		sourceTemplateService: sourceTemplateService,
	}
}

func (con *SourceTemplateController) AddSourceTemplate(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Request.Header.Get("User-Id"))
	// Parse request body
	jsonBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	template := entity.SourceTemplate{}
	err = json.Unmarshal(jsonBytes, &template)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := utils.GetValueOnlyRequestContext(c)

	id, err := con.sourceTemplateService.AddSourceTemplate(ctx, template, int64(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (con *SourceTemplateController) GetSourceTemplateByType(c *gin.Context) {
	// Get template type from query parameter
	templateType := c.Param("type")

	ctx := utils.GetValueOnlyRequestContext(c)

	template, err := con.sourceTemplateService.GetSourceTemplateByType(ctx, templateType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, template)
}
