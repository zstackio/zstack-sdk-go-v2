// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateCertificate creates Certificate
func (cli *ZSClient) CreateCertificate(params param.CreateCertificateParam) (*view.CreateCertificateEventView, error) {
	resp := view.CreateCertificateEventView{}
	if err := cli.Post("v1/certificates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
