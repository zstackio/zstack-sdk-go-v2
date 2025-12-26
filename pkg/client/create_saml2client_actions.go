// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSAML2Client creates SAML2Client
func (cli *ZSClient) CreateSAML2Client(params param.CreateSAML2ClientParam) (*view.CreateSAML2ClientEventView, error) {
	resp := view.CreateSAML2ClientEventView{}
	if err := cli.Post("v1/create/saml2/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
