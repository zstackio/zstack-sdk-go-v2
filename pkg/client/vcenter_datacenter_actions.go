// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVCenterDatacenter queries VCenterDatacenter list
func (cli *ZSClient) QueryVCenterDatacenter(params *param.QueryParam) ([]view.VCenterDatacenterInventoryView, error) {
	var resp []view.VCenterDatacenterInventoryView
	return resp, cli.List("v1/vcenters/datacenters", params, &resp)
}
