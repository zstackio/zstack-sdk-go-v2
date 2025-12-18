// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetTrashOnPrimaryStorage 获取TrashOnPrimaryStorage详情
func (cli *ZSClient) GetTrashOnPrimaryStorage(uuid string) (*view.GetTrashOnPrimaryStorageView, error) {
	var resp view.GetTrashOnPrimaryStorageView
	if err := cli.Get("v1/primary-storage/trash", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

