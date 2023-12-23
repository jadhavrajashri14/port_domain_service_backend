package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
	"port_domain_service_backend/internal/core/domain"

	"github.com/sirupsen/logrus"
)

type PortDomainsRepository struct {
	PortDomains map[string]map[string]interface{}
}

func readJson() (map[string]interface{}, error) { //nolint
	logrus.Info(os.Getwd())

	p, err := filepath.Abs("/home/rajashrijadhav/RajashriJadhavData/RAJASHRI_ASSIGNMENTS/port_domain_service_backend/testdata/ports.json")
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	logrus.Info(p)

	data, err := os.ReadFile(p)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	logrus.Info("Successfully opened ports json file")

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return m, nil
}

func NewPortDomainsRepository() *PortDomainsRepository {
	return &PortDomainsRepository{
		PortDomains: make(map[string]map[string]interface{}),
	}
}

func (pdr *PortDomainsRepository) CreatePortDomain(fileBytes []byte) error {
	var pds map[string]interface{}
	err := json.Unmarshal(fileBytes, &pds)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	counter := 0
	for port, details := range pds {
		pDetails := details.(map[string]interface{})
		pdr.PortDomains[port] = pDetails
		counter++
	}
	logrus.Infof("%d : json file data stored in memory map", counter)

	/**for k, _ := range pdr.PortDomains {
		logrus.Info(k)
	}
	**/
	return nil
}

func (pdr *PortDomainsRepository) UpdatePortDomain(pd domain.PortDetails) (domain.PortDetail, error) {
	logrus.Info("PortDomainsRepository UpdatePortDomain called!")

	for k, details := range pd {
		logrus.Info(k)
		if val, ok := pdr.PortDomains[k]; ok {
			logrus.Info("key present in original map")
			logrus.Info(val)
			pdr.PortDomains[k] = details
			pd := details
			return pd, nil
		} else {
			logrus.Info("key not present in original map, can't update")
			return nil, nil
		}
	}
	return nil, nil
}
