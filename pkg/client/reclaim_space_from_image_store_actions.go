// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReclaimSpaceFromImageStore 操作ReclaimSpaceFromImageStore
func (cli *ZSClient) ReclaimSpaceFromImageStore(uuid string, params param.ReclaimSpaceFromImageStoreParam) (*view.ReclaimSpaceFromImageStoreEventView, error) {
	resp := view.ReclaimSpaceFromImageStoreEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

