package service

import (
	"github.com/aluka-7/api-base/app"
	"github.com/aluka-7/datasource/base"
	"github.com/aluka-7/datasource/search"
	"xorm.io/xorm"
)

func NewCompanyRepository(orm *xorm.Engine) app.ICompanyRepository {
	return &companyRepository{
		base.NewBaseRepository(orm, nil),
		map[string]search.Filter{},
	}
}

type companyRepository struct {
	base.BaseRepository
	column map[string]search.Filter
}
