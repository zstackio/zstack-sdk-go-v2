// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddCCSCertificate adds CCSCertificate
func (cli *ZSClient) AddCCSCertificate(params param.AddCCSCertificateParam) (*view.AddCCSCertificateEventView, error) {
	resp := view.AddCCSCertificateEventView{}
	if err := cli.Post("v1/crypto/ccs-certificate/add", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
