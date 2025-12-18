// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectPrimaryStorage 操作ReconnectPrimaryStorage
func (cli *ZSClient) ReconnectPrimaryStorage(uuid string, params param.ReconnectPrimaryStorageParam) (*view.ReconnectPrimaryStorageEventView, error) {
	resp := view.ReconnectPrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

