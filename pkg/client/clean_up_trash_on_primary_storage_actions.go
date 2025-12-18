// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanUpTrashOnPrimaryStorage 操作CleanUpTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpTrashOnPrimaryStorage(uuid string, params param.CleanUpTrashOnPrimaryStorageParam) (*view.CleanUpTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpTrashOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/trash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

