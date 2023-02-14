//go:build wireinject
// +build wireinject

//go:generate wire

package wire

import (
	"github.com/aluka-7/api-base/app"
	ar "github.com/aluka-7/api-base/app/repository"
	as "github.com/aluka-7/api-base/app/service"
	"github.com/google/wire"
	"xorm.io/xorm"
)

const (
	SystemId = "1000"
)

func InitializeCompanyService(*xorm.Engine) app.ICompanyService {
	panic(wire.Build(as.NewCompanyService, ar.NewCompanyRepository))
}
