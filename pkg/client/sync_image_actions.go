// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncImage operates on SyncImage
func (cli *ZSClient) SyncImage(uuid string, params param.SyncImageParam) (*view.SyncImageEventView, error) {
	resp := view.SyncImageEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
