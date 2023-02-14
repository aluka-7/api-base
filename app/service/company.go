package service

import (
	"github.com/aluka-7/api-base/app"
)

func NewCompanyService(cr app.ICompanyRepository) app.ICompanyService {
	return &companyService{
		cr,
	}
}

type companyService struct {
	cr app.ICompanyRepository
}
