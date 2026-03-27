// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryKvmHypervisorInfo queries KvmHypervisorInfo list
func (cli *ZSClient) QueryKvmHypervisorInfo(ctx context.Context, params *param.QueryParam) ([]view.KvmHypervisorInfoInventoryView, error) {
	var resp []view.KvmHypervisorInfoInventoryView
	return resp, cli.List(ctx, "v1/hosts/kvm/hypervisor/info", params, &resp)
}

func (cli *ZSClient) GetKvmHypervisorInfo(ctx context.Context, uuid string) (*view.KvmHypervisorInfoInventoryView, error) {
	var resp view.KvmHypervisorInfoInventoryView
	if err := cli.Get(ctx, "v1/hosts/kvm/hypervisor/info", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageKvmHypervisorInfo Pagination
func (cli *ZSClient) PageKvmHypervisorInfo(ctx context.Context, params *param.QueryParam) ([]view.KvmHypervisorInfoInventoryView, int, error) {
	var kvmHypervisorInfos []view.KvmHypervisorInfoInventoryView
	total, err := cli.Page(ctx, "v1/hosts/kvm/hypervisor/info", params, &kvmHypervisorInfos)
	return kvmHypervisorInfos, total, err
}
