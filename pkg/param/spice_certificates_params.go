// Copyright (c) ZStack.io, Inc.

package param

// GetSpiceCertificatesDetailParam GetSpiceCertificates详细参数
type GetSpiceCertificatesDetailParam struct {
}

// GetSpiceCertificatesParam GetSpiceCertificates请求参数
type GetSpiceCertificatesParam struct {
	BaseParam
	Params GetSpiceCertificatesDetailParam `json:"params"` // 详细参数
}

