// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddExternalPrimaryStorage adds ExternalPrimaryStorage
func (cli *ZSClient) AddExternalPrimaryStorage(ctx context.Context, params param.AddExternalPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/primary-storage/addon", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateExternalPrimaryStorage updates ExternalPrimaryStorage
func (cli *ZSClient) UpdateExternalPrimaryStorage(ctx context.Context, uuid string, params param.UpdateExternalPrimaryStorageParam) (*view.ExternalPrimaryStorageInventoryView, error) {
	resp := view.ExternalPrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage/addon", uuid, "", map[string]interface{}{
		"updateExternalPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
