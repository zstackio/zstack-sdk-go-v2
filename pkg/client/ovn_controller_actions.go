// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryOvnController queries OvnController list
func (cli *ZSClient) QueryOvnController(params *param.QueryParam) ([]view.OvnControllerInventoryView, error) {
	var resp []view.OvnControllerInventoryView
	return resp, cli.List("v1/ovn-controllers", params, &resp)
}

func (cli *ZSClient) GetOvnController(uuid string) (*view.OvnControllerInventoryView, error) {
	var resp view.OvnControllerInventoryView
	if err := cli.Get("v1/ovn-controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
