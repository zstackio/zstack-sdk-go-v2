// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSSOClient deletes SSOClient
func (cli *ZSClient) DeleteSSOClient(params param.DeleteSSOClientParam) (*view.DeleteSSOClientEventView, error) {
	resp := view.DeleteSSOClientEventView{}
	if err := cli.PostWithRespKey("v1/delete/sso/client", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSSOClient gets SSOClient
func (cli *ZSClient) GetSSOClient() (*view.GetSSOClientView, error) {
	var resp view.GetSSOClientView
	if err := cli.GetWithRespKey("v1/get/sso/client", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
