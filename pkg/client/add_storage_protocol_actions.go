// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddStorageProtocol adds StorageProtocol
func (cli *ZSClient) AddStorageProtocol(params param.AddStorageProtocolParam) (*view.AddStorageProtocolEventView, error) {
	resp := view.AddStorageProtocolEventView{}
	if err := cli.Post("v1/primary-storage/protocol", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
