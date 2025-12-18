// Copyright (c) ZStack.io, Inc.

package param

// CreateSlbGroupDetailParam CreateSlbGroup detail param
type CreateSlbGroupDetailParam struct {
	Name string `json:"name" validate:"required"`
	SlbOfferingUuid string `json:"slbOfferingUuid" validate:"required"`
	FrontEndL3NetworkUuid string `json:"frontEndL3NetworkUuid" validate:"required"`
	BackendL3NetworkUuids []string `json:"backendL3NetworkUuids,omitempty"`
	BackendType string `json:"backendType,omitempty"`
	DeployType string `json:"deployType,omitempty"`
	Description string `json:"description,omitempty"`
	MonitorIps []string `json:"monitorIps,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSlbGroupParam CreateSlbGroup request param
type CreateSlbGroupParam struct {
	BaseParam
	Params CreateSlbGroupDetailParam `json:"params"`
}
