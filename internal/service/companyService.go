package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mauriciodm1998/CompanyService/internal/canonical"
	"github.com/mauriciodm1998/CompanyService/internal/repositories"
)

type Service interface {
	Get(ctx context.Context, id string) (*canonical.Company, error)
	Create(ctx context.Context, company canonical.Company) (string, error)
}

type companyService struct {
	repo repositories.Repository
}

func New() Service {
	return &companyService{
		repo: repositories.New(),
	}
}

func (s companyService) Get(ctx context.Context, id string) (*canonical.Company, error) {
	company, err := s.repo.Get(ctx, id)
	if err != nil {
		return company, err
	}

	return company, nil
}

func (s companyService) Create(ctx context.Context, company canonical.Company) (string, error) {
	if company.Id == "" {
		company.Id = uuid.NewString()
	}

	id, err := s.repo.Create(ctx, company)
	if err != nil {
		return "", err
	}

	return id, nil
}
