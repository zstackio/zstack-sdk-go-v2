// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReclaimSpaceFromImageStore operates on ReclaimSpaceFromImageStore
func (cli *ZSClient) ReclaimSpaceFromImageStore(uuid string, params param.ReclaimSpaceFromImageStoreParam) (*view.ReclaimSpaceFromImageStoreEventView, error) {
	resp := view.ReclaimSpaceFromImageStoreEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
