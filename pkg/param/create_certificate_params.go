// Copyright (c) ZStack.io, Inc.

package param

// CreateCertificateDetailParam CreateCertificate detail param
type CreateCertificateDetailParam struct {
	Name string `json:"name" validate:"required"`
	Certificate string `json:"certificate" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCertificateParam CreateCertificate request param
type CreateCertificateParam struct {
	BaseParam
	Params CreateCertificateDetailParam `json:"params"`
}
