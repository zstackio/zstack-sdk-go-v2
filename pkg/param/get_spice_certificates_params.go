// Copyright (c) ZStack.io, Inc.

package param

// GetSpiceCertificatesDetailParam GetSpiceCertificates detail param
type GetSpiceCertificatesDetailParam struct {
}

// GetSpiceCertificatesParam GetSpiceCertificates request param
type GetSpiceCertificatesParam struct {
	BaseParam
	Params GetSpiceCertificatesDetailParam `json:"params"`
}
