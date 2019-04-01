package controllers

import (
	"github.com/dmjones/revel-interceptor-issue/app"
	"github.com/revel/revel"
)

type BaseController struct {
	*revel.Controller
}

func (c *BaseController) Before() (revel.Result, *BaseController) {
	revel.AppLog.Info("Base Before")

	if app.Switch {
		return c.Redirect("http://www.google.com"), c
	}

	app.Switch = !app.Switch

	return nil, c
}

type ExtendedController struct {
	BaseController
}

func (c *ExtendedController) Before() (revel.Result, *ExtendedController) {
	revel.AppLog.Info("Extended Before")
	return nil, c
}
