// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryKvmHypervisorInfo queries KvmHypervisorInfo list
func (cli *ZSClient) QueryKvmHypervisorInfo(params *param.QueryParam) ([]view.KvmHypervisorInfoInventoryView, error) {
	var resp []view.KvmHypervisorInfoInventoryView
	return resp, cli.List("v1/hosts/kvm/hypervisor/info", params, &resp)
}
