// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVpcSnatState queries VpcSnatState list
func (cli *ZSClient) QueryVpcSnatState(params *param.QueryParam) ([]view.VpcSnatStateInventoryView, error) {
	var resp []view.VpcSnatStateInventoryView
	return resp, cli.List("v1/vpc/virtual-routers/networkservicestate/snat", params, &resp)
}

func (cli *ZSClient) GetVpcSnatState(uuid string) (*view.VpcSnatStateInventoryView, error) {
	var resp view.VpcSnatStateInventoryView
	if err := cli.Get("v1/vpc/virtual-routers/networkservicestate/snat", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
