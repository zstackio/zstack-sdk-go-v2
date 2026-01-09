// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateSAML2Client updates SAML2Client
func (cli *ZSClient) UpdateSAML2Client(uuid string, params param.UpdateSAML2ClientParam) (*view.SAML2ClientInventoryView, error) {
	var resp view.UpdateSAML2ClientEventView
	if err := cli.Put("v1/update/saml2/client", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateSAML2Client creates SAML2Client
func (cli *ZSClient) CreateSAML2Client(params param.CreateSAML2ClientParam) (*view.SAML2ClientInventoryView, error) {
	var resp view.CreateSAML2ClientEventView
	if err := cli.Post("v1/create/saml2/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
