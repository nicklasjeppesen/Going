package controller

type HomeController struct {
	ControllerBase
}

func (c *HomeController) Home() Result {
	return View("index", Params{"Title": "Going App", "Message": "Welcome to Going"})

}
