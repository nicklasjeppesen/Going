package routes

import (
	//	. "myapp/internal/app/http/controller"

	web "github.com/nicklasjeppesen/going_internal/super/customrouter"
)

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
|
*/

func Samplerouter() *web.MyRouter {

	SampleRouter := web.NewMyRouter()
	//var sampleController = SampleController{}

	/*
		SampleRouter.GET("/nicklas", func(w http.ResponseWriter, r *http.Request) {

			fmt.Fprint(w, "passed")
		}).AddMiddleware(middleware.MiddlewareCors)

		// Single example with Specefic route middleware

		SampleRouter.GET("/", sampleController.Get)

		SampleRouter.GET("/sample/index", sampleController.ShowUser).Name("sample.index")
		SampleRouter.POST("/store", sampleController.Store).Name("sample.store")

		SampleRouter.GET("/sample/show/{id}", sampleController.ShowUser).Name("sample.ShowInput")

		SampleRouter.GET("/hejverden/{id}/{name}", sampleController.APIONE).
			AddMiddleware(middleware.MiddlewareFour).
			AddMiddleware(middleware.MiddlewareFive)

		// -------------------- API -------------------//
		SampleRouter.Groups("/api",
			SampleRouter.GET("/one", sampleController),
			SampleRouter.POST("/two", sampleController.APIONE),
		)

		SampleRouter.GroupsWithMiddleware("/api2", middleware.AuthMiddleware,
			SampleRouter.GET("/three", sampleController.APIONE).AddMiddleware(middleware.LoggingMiddleware),
			SampleRouter.GET("/four", sampleController.APIONE),
		)
	*/
	return SampleRouter

}
