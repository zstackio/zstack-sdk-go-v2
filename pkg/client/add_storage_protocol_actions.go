// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddStorageProtocol 操作AddStorageProtocol
func (cli *ZSClient) AddStorageProtocol(params param.AddStorageProtocolParam) (*view.AddStorageProtocolEventView, error) {
	resp := view.AddStorageProtocolEventView{}
	if err := cli.Post("v1/primary-storage/protocol", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

