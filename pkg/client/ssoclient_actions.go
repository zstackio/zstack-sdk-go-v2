// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSSOClient deletes SSOClient
func (cli *ZSClient) DeleteSSOClient(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/delete/sso/client", uuid, string(deleteMode))
}
// GetSSOClient gets SSOClient by uuid
func (cli *ZSClient) GetSSOClient(uuid string) (*view.SSOClientInventoryView, error) {
	var resp view.SSOClientInventoryView
	if err := cli.Get("v1/get/sso/client", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
