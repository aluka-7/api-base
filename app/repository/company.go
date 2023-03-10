package service

import (
	"github.com/aluka-7/api-base/app"
	"github.com/aluka-7/cache"
	dbase "github.com/aluka-7/datasource/base"
	"github.com/aluka-7/datasource/search"
	mbase "github.com/aluka-7/gomongo/base"
	"go.mongodb.org/mongo-driver/mongo"
	"xorm.io/xorm"
)

func NewCompanyRepository(orm *xorm.Engine, gmc *mongo.Database, ce cache.Provider) app.ICompanyRepository {
	return &companyRepository{
		dbase.NewBaseRepository(orm, nil),
		mbase.NewBaseRepository(gmc, nil),
		map[string]search.Filter{},
		ce,
	}
}

type companyRepository struct {
	dbase  dbase.BaseRepository
	mbase  mbase.BaseRepository
	column map[string]search.Filter
	ce     cache.Provider
}
