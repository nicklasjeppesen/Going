package controller

type AdminController struct {
	ControllerBase
}

func (c *AdminController) Home() Result {
	return Response.Print("This is admin page")

}

func (c *AdminController) Users() Result {
	return Response.Print("This is Users page")
}

func (c *AdminController) Company() Result {
	return Response.Print("This is Companies page")
}
