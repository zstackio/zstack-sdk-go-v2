// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryCephOsdGroup queries CephOsdGroup list
func (cli *ZSClient) QueryCephOsdGroup(params *param.QueryParam) ([]view.CephOsdGroupInventoryView, error) {
	var resp []view.CephOsdGroupInventoryView
	return resp, cli.List("v1/primary-storage/ceph/osdgroups", params, &resp)
}
