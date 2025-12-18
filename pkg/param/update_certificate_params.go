// Copyright (c) ZStack.io, Inc.

package param

// UpdateCertificateDetailParam UpdateCertificate detail param
type UpdateCertificateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateCertificateParam UpdateCertificate request param
type UpdateCertificateParam struct {
	BaseParam
	Params UpdateCertificateDetailParam `json:"params"`
}
