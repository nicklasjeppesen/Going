package docs

import (
	"net/http"

	util "github.com/nicklasjeppesen/going_internal/super/util"

	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8443/
// @BasePath

func RegisterSwagger(router *http.ServeMux) {
	var url = util.GetURL()
	router.HandleFunc("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL(url+"/swagger/doc.json"), //The url pointing to API definition
	))

}
