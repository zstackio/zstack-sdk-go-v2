// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedResource queries SharedResource list
func (cli *ZSClient) QuerySharedResource(params *param.QueryParam) ([]view.SharedResourceInventoryView, error) {
	var resp []view.SharedResourceInventoryView
	return resp, cli.List("v1/accounts/resources", params, &resp)
}

func (cli *ZSClient) GetSharedResource(uuid string) (*view.SharedResourceInventoryView, error) {
	var resp view.SharedResourceInventoryView
	if err := cli.Get("v1/accounts/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
