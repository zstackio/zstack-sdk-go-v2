// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// MoveDirectory operates on MoveDirectory
func (cli *ZSClient) MoveDirectory(uuid string, params param.MoveDirectoryParam) (*view.MoveDirectoryEventView, error) {
	resp := view.MoveDirectoryEventView{}
	if err := cli.Put("v1/move/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
