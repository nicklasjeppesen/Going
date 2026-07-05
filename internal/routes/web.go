package routes

import (
	//. "myapp/internal/app/http/controller"

	. "myapp/internal/app/http/controller"
	. "myapp/internal/app/http/controller/auth"
	"myapp/internal/app/providers/container"

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
func Webrouter() *web.MyRouter {

	webrouter := web.NewMyRouter().UseContainer(container.GetContainer())

	var homeController = new(HomeController)
	var registerController = RegisterController{}
	var loginController = LoginController{}

	//webrouter.GET("/", homeController.Home).Name("home.front")
	webrouter.Get("/", homeController, "Home").Name("home.front")

	webrouter.Get("/register", registerController.RegisterGet).Name("auth.register")
	webrouter.Post("/registerPost", registerController.Register).Name("auth.register.post")

	webrouter.Post("/login", loginController.Login)
	webrouter.Get("/login", loginController.LoginGet).Name("auth.login")
	webrouter.Post("/logout", loginController.Logout)

	webrouter.Get("/protected", loginController.Protected).AddMiddleware(middleware.JWTMiddleware)

	return webrouter
}
