// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddLocalPrimaryStorage 操作AddLocalPrimaryStorage
func (cli *ZSClient) AddLocalPrimaryStorage(params param.AddLocalPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/local-storage", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

