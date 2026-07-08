package routes

import (
	//. "myapp/internal/app/http/controller"

	"fmt"
	. "myapp/internal/app/http/controller"
	. "myapp/internal/app/http/controller/auth"
	"net/http"

	web "github.com/nicklasjeppesen/going_internal/super/customrouter"
	"github.com/nicklasjeppesen/going_internal/super/middleware"
)

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
| (Text is stolen from laravel frameworks, because laravel is also a great framework)
|
*/
func Webrouter(webrouter *web.MyRouter) *web.MyRouter {

	var homeController = new(HomeController)
	var registerController = RegisterController{}
	var loginController = LoginController{}
	var sampleController = SampleController{}

	// 404 - page not found
	webrouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "page not found :(")
	})

	// Home page
	webrouter.Get("/{$}", homeController, "Home").Name("home.front")

	webrouter.Post("/tester", sampleController.Tester)

	webrouter.Get("/register", registerController.RegisterGet).Name("auth.register")
	webrouter.Post("/registerPost", registerController.Register).Name("auth.register.post")

	webrouter.Post("/login", loginController.Login)
	webrouter.Get("/login", loginController.LoginGet).Name("auth.login")
	webrouter.Post("/logout", loginController.Logout)

	webrouter.Get("/protected", loginController.Protected).AddMiddleware(middleware.JWTMiddleware)

	// Admin
	adminRouter := webrouter.Group("/admin")

	adminRouterUsers := adminRouter.Group("/users")
	adminRoutercompany := adminRouter.Group("/company")

	adminRouter.Get("/index", new(AdminController), "Home")
	adminRouterUsers.Get("/index", new(AdminController), "Users")
	adminRoutercompany.Get("/index", new(AdminController), "Company")

	return webrouter
}
