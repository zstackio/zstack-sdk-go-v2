// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanUpStorageTrashOnPrimaryStorage 操作CleanUpStorageTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpStorageTrashOnPrimaryStorage(uuid string, params param.CleanUpStorageTrashOnPrimaryStorageParam) (*view.CleanUpStorageTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpStorageTrashOnPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/storagetrash/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

