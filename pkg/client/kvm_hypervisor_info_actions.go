// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryKvmHypervisorInfo queries KvmHypervisorInfo list
func (cli *ZSClient) QueryKvmHypervisorInfo(params *param.QueryParam) ([]view.KvmHypervisorInfoInventoryView, error) {
	var resp []view.KvmHypervisorInfoInventoryView
	return resp, cli.List("v1/hosts/kvm/hypervisor/info", params, &resp)
}

func (cli *ZSClient) GetKvmHypervisorInfo(uuid string) (*view.KvmHypervisorInfoInventoryView, error) {
	var resp view.KvmHypervisorInfoInventoryView
	if err := cli.Get("v1/hosts/kvm/hypervisor/info", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
