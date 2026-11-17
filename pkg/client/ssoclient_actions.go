// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSSOClient deletes SSOClient
func (cli *ZSClient) DeleteSSOClient(params param.DeleteSSOClientParam) (*view.SSOClientInventoryView, error) {
	resp := view.SSOClientInventoryView{}
	if err := cli.Post("v1/delete/sso/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// GetSSOClient gets SSOClient by uuid
func (cli *ZSClient) GetSSOClient() (*view.SSOClientInventoryView, error) {
	var resp view.SSOClientInventoryView
	if err := cli.GetWithRespKey("v1/get/sso/client", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SsoClientPushData operates on SsoClientPushData
func (cli *ZSClient) SsoClientPushData(params param.SsoClientPushDataParam) (*view.SsoClientPushDataEventView, error) {
	resp := view.SsoClientPushDataEventView{}
	if err := cli.PutWithRespKey("v1/sso/resource/data/push", "", "", map[string]interface{}{
		"ssoClientPushData": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
