// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCasClient creates CasClient
func (cli *ZSClient) CreateCasClient(params param.CreateCasClientParam) (*view.CasClientInventoryView, error) {
	resp := view.CasClientInventoryView{}
	if err := cli.Post("v1/create/cas/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateCasClient updates CasClient
func (cli *ZSClient) UpdateCasClient(params param.UpdateCasClientParam) (*view.CasClientInventoryView, error) {
	resp := view.CasClientInventoryView{}
	if err := cli.Post("v1/update/cas/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOAuthClient creates OAuthClient
func (cli *ZSClient) CreateOAuthClient(params param.CreateOAuthClientParam) (*view.OAuth2ClientInventoryView, error) {
	resp := view.OAuth2ClientInventoryView{}
	if err := cli.Post("v1/create/oauth2/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOAuthClient updates OAuthClient
func (cli *ZSClient) UpdateOAuthClient(params param.UpdateOAuthClientParam) (*view.OAuth2ClientInventoryView, error) {
	resp := view.OAuth2ClientInventoryView{}
	if err := cli.Post("v1/update/oauth2/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
