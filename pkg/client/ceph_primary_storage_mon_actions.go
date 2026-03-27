// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateCephPrimaryStorageMon updates CephPrimaryStorageMon
func (cli *ZSClient) UpdateCephPrimaryStorageMon(ctx context.Context, monUuid string, params param.UpdateCephPrimaryStorageMonParam) (*view.CephPrimaryStorageInventoryView, error) {
	resp := view.CephPrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage/ceph/mons", monUuid, "", map[string]interface{}{
		"updateCephPrimaryStorageMon": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
