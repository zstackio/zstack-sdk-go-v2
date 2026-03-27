// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddExternalBackupStorage adds ExternalBackupStorage
func (cli *ZSClient) AddExternalBackupStorage(ctx context.Context, params param.AddExternalBackupStorageParam) (*view.ExternalBackupStorageInventoryView, error) {
	resp := view.ExternalBackupStorageInventoryView{}
	if err := cli.Post(ctx, "v1/backup-storage/addon", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
