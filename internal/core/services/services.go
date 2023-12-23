package services

import (
	"fmt"
	"port_domain_service_backend/internal/adapters/repository"
	"port_domain_service_backend/internal/core/domain"

	"github.com/sirupsen/logrus"
)

type PortDomainsService struct {
	repo repository.PortDomainsRepository
}

func NewPortDomainsService(repo repository.PortDomainsRepository) *PortDomainsService {
	return &PortDomainsService{
		repo: repo,
	}
}

func (pds *PortDomainsService) CreatePortDomain(fileBytes []byte) error {
	logrus.Info("PortDomainsService CreatePortDomain called!")
	err := pds.repo.CreatePortDomain(fileBytes)
	if err != nil {
		return err
	}
	return nil
}

func (pds *PortDomainsService) UpdatePortDomain(pd domain.PortDetails) (domain.PortDetail, error) {
	details, err := pds.repo.UpdatePortDomain(pd)
	if err != nil {
		return nil, fmt.Errorf("unable to update port domain : %s", err.Error())
	}
	return details, nil
}
