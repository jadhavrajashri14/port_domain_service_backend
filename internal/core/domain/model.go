package domain

type PortDomain struct {
	PortID     string                 `json:"port_id"`
	PortDetail map[string]interface{} `json:"port_detail"`
}

type PortDetail map[string]interface{}

type PortDetails map[string]map[string]interface{}
