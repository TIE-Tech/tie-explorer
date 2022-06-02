package controllers

type DefaultController struct {
	BaseController
}

func (o *DefaultController) Default() {
	o.Ctx.WriteString("Tie Explorer API Server")
}
