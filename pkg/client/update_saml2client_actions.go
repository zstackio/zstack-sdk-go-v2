// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSAML2Client updates SAML2Client
func (cli *ZSClient) UpdateSAML2Client(uuid string, params param.UpdateSAML2ClientParam) (*view.UpdateSAML2ClientEventView, error) {
	resp := view.UpdateSAML2ClientEventView{}
	if err := cli.Put("v1/update/saml2/client", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
