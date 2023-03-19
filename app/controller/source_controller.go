package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rudderstack_source_app/entity"
	"github.com/rudderstack_source_app/services/source"
	"github.com/rudderstack_source_app/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SourceController struct {
	sourceService source.SourceServiceInterface
}

func NewSourceController(sourceService source.SourceServiceInterface) *SourceController {
	return &SourceController{
		sourceService: sourceService,
	}
}

func (con *SourceController) CreateSource(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Request.Header.Get("User-Id"))

	jsonBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	source := entity.Source{}
	err = json.Unmarshal(jsonBytes, &source)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx := utils.GetValueOnlyRequestContext(c)

	err = con.sourceService.CreateSource(ctx, source, int64(userId))
	if err != nil {

	}
	c.JSON(http.StatusOK, gin.H{"message": "Source created successfully"})
}

func (con *SourceController) GetAllSources(c *gin.Context) {
	sourceType := c.Param("type")
	ctx := utils.GetValueOnlyRequestContext(c)

	sources, err := con.sourceService.GetAllSources(ctx, sourceType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *sources)
}
