package controller

import (
	"github.com/aluka-7/api-base/app"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CompanyController struct {
	cs app.ICompanyService
}

func NewCompanyController(eng *echo.Group, prefix string, cs app.ICompanyService) {
	ctrl := CompanyController{cs: cs}
	group := eng.Group(prefix)
	group.GET("", ctrl.Index)
}

func (ctrl CompanyController) Index(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello world")
}
