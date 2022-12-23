package main

import (
	"net/http"

	// load docs for swagger UI
	_ "github.com/itsprdp/swaggo-custom-types/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type sampleResponse struct {
	Data CustomType `json:"data"`
}

type field string
type subType struct {
	Fields []field `json:"fields"`
}

type CustomType struct {
	CustomField subType `json:"customField" swaggertype:"object,main.subType"`
}

// ShowCustomType godoc
//
//	@Summary		Sample response
//	@Description	returns a response with custom struct
//	@Tags			 custom_struct
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	sampleResponse
//	@Router			/custom_type [get]
func ShowCustomType(ctx *gin.Context) {
	response := sampleResponse{}
	ctx.JSON(http.StatusOK, response)
}

// Following comment lines are meant for the swaggo and gin-swagger to auto
// generate the swagger spec files.
//
// @title                       Swaggo API
// @version                     1.0
// @description                 Sample API App to reproduce the bug
// @host                        localhost:8000
// @BasePath                    /
// @schemes                     http
func main() {
	router := gin.Default()

	// health endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// swagger docs page as index page
	router.GET("/", func(c *gin.Context) {
		docsURL := "/swagger/index.html"

		c.Redirect(http.StatusFound, docsURL)
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/custom_type", ShowCustomType)

	router.Run("0.0.0.0:8000")
}
