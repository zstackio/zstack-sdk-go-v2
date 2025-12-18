// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshFiberChannelStorage 操作RefreshFiberChannelStorage
func (cli *ZSClient) RefreshFiberChannelStorage(params param.RefreshFiberChannelStorageParam) (*view.RefreshFiberChannelStorageEventView, error) {
	resp := view.RefreshFiberChannelStorageEventView{}
	if err := cli.Post("v1/storage-devices/fiber-channel/controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

