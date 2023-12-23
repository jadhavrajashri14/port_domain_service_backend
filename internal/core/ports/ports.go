package ports

import (
	"port_domain_service_backend/internal/core/domain"
)

type PortDomainsService interface {
	CreatePortDomain(data []byte) error
	UpdatePortDomain(portDomain domain.PortDetails) (*domain.PortDetail, error)
}

type PortDomainRepository interface {
	CreatePortDomain(data []byte) error
	UpdatePortDomain(portDomains domain.PortDetails) (*domain.PortDetail, error)
}
