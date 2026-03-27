// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateCephBackupStorageMon updates CephBackupStorageMon
func (cli *ZSClient) UpdateCephBackupStorageMon(ctx context.Context, monUuid string, params param.UpdateCephBackupStorageMonParam) (*view.CephBackupStorageInventoryView, error) {
	resp := view.CephBackupStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/ceph/mons", monUuid, "", map[string]interface{}{
		"updateCephBackupStorageMon": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
