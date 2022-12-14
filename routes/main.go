package routes

import (
	"encoding/json"
	userRoute "gin-mongo-boilerplate/src/user/routes"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/gin-gonic/gin"
	. "github.com/shipu/artifact"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Register() {
	BaseRoute()
	SwaggerRoute()

	userRoute.UserSetup()
}

func BaseRoute() {
	// Example Route
	Router.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"app": Config.GetString("App.Name"),
		}

		//or
		//data := gin.H{
		//	"message": "Hello World",
		//}

		Res.Code(200).
			Message("success").
			Data(data).Json(c)
	})
}

func SwaggerRoute() {
	Router.GET("swagger/swagger.json", func(c *gin.Context) {
		f, err := os.Open("./docs/swagger.json")
		if err != nil {
			Res.Code(200).
				Message("err").
				Data(err).
				AbortWithStatusJSON(c)
			return
		}
		defer f.Close()

		dec := json.NewDecoder(f)
		data := make(map[string]interface{}, 0)
		if err := dec.Decode(&data); err != nil {
			Res.Code(200).
				Message("err").
				Data(err).
				AbortWithStatusJSON(c)
			return
		}
		appUrl := Config.GetString("App.Url")
		appUrl = strings.Replace(appUrl, "https", "", -1)
		appUrl = strings.Replace(appUrl, "http", "", -1)
		data["host"] = appUrl

		Res.
			Raw(data).
			Json(c)
	})

	Router.GET("swagger/swagger.yaml", func(c *gin.Context) {
		f, err := ioutil.ReadFile("./docs/swagger.yaml")
		if err != nil {
			Res.Code(200).
				Message("err").
				Data(err).
				AbortWithStatusJSON(c)
			return
		}

		data := make(map[string]interface{}, 0)
		err = yaml.Unmarshal([]byte(f), &data)
		if err != nil {
			Res.Code(200).
				Message("err").
				Data(err).
				AbortWithStatusJSON(c)
			return
		}
		appUrl := Config.GetString("App.Url")
		appUrl = strings.Replace(appUrl, "https://", "", -1)
		appUrl = strings.Replace(appUrl, "http://", "", -1)
		data["host"] = appUrl

		Res.
			Raw(data).
			Yaml(c)
	})

	Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(Config.GetString("App.Url")+"/swagger/swagger.yaml"), ginSwagger.DefaultModelsExpandDepth(-1)))
}
