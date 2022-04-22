package repositories

import (
	"context"
	"encoding/json"

	"github.com/mauriciodm1998/CompanyService/internal/canonical"
	"github.com/mauriciodm1998/pkg/abkv"

	"github.com/MurilloVaz/bitcask"
	"github.com/mauriciodm1998/internal/config"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Get(ctx context.Context, id string) (*canonical.Company, error)
	Create(ctx context.Context, company canonical.Company) (string, error)
}

type companyRepository struct {
	db *bitcask.Bitcask
}

func New() Repository {
	db, err := abkv.Open("companyService", config.DbPath)

	if err != nil {
		logrus.WithError(err).Fatal("Cannot open DB")
	}

	return &companyRepository{db}
}

func (r *companyRepository) Get(ctx context.Context, id string) (*canonical.Company, error) {
	company, err := r.db.Get([]byte(id))

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (r *companyRepository) Create(ctx context.Context, company canonical.Company) (string, error) {
	companyByte, err := json.Marshal(company)
	if err != nil {
		return "", err
	}

	err = r.db.Put([]byte(company.Id), companyByte)
	if err != nil {
		return "", err
	}

	return company.Id, nil
}
