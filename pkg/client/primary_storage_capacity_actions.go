// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// SyncPrimaryStorageCapacity operates on PrimaryStorageCapacity
func (cli *ZSClient) SyncPrimaryStorageCapacity(ctx context.Context, primaryStorageUuid string, params param.SyncPrimaryStorageCapacityParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage", primaryStorageUuid, "", map[string]interface{}{
		"syncPrimaryStorageCapacity": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// GetPrimaryStorageCapacity gets PrimaryStorageCapacity by uuid
func (cli *ZSClient) GetPrimaryStorageCapacity(ctx context.Context) (*view.GetPrimaryStorageCapacityView, error) {
	var resp view.GetPrimaryStorageCapacityView
	if err := cli.GetWithRespKey(ctx, "v1/primary-storage/capacities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
