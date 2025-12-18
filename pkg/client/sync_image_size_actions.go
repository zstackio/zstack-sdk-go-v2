// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncImageSize 操作SyncImageSize
func (cli *ZSClient) SyncImageSize(uuid string, params param.SyncImageSizeParam) (*view.SyncImageSizeEventView, error) {
	resp := view.SyncImageSizeEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

