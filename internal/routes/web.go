package routes

import (
	//. "myapp/internal/app/http/controller"

	. "myapp/internal/app/http/controller"
	. "myapp/internal/app/http/controller/auth"

	web "github.com/nicklasjeppesen/going_internal/super/customweb"
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

	webrouter := web.NewMyRouter()
	var loginController = LoginController{}
	var homeController = HomeController{}

	webrouter.GET("/home", homeController.Home).Name("home.home")
	webrouter.GET("/login", loginController.LoginGet).Name("auth.login")

	/*
		var homeController = HomeController{}
		var registerController = RegisterController{}
		var loginController = LoginController{}

		webrouter.GET("/register", registerController.RegisterGet).Name("auth.register")
		webrouter.POST("/register", registerController.Register)

		webrouter.POST("/login", loginController.Login)
		webrouter.GET("/login", loginController.LoginGet).Name("auth.login")
		webrouter.GET("/logout", loginController.Logout)

		webrouter.GET("/protected", loginController.Protected).AddMiddleware(middleware.JWTMiddleware)
	*/
	return webrouter
}
