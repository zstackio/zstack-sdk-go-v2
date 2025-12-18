// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncImage operates on SyncImage
func (cli *ZSClient) SyncImage(uuid string, params param.SyncImageParam) (*view.SyncImageEventView, error) {
	resp := view.SyncImageEventView{}
	if err := cli.Put("v1/backup-storage/image-store/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
