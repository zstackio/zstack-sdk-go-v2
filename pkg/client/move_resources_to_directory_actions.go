// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// MoveResourcesToDirectory operates on MoveResourcesToDirectory
func (cli *ZSClient) MoveResourcesToDirectory(uuid string, params param.MoveResourcesToDirectoryParam) (*view.MoveResourcesToDirectoryEventView, error) {
	resp := view.MoveResourcesToDirectoryEventView{}
	if err := cli.Put("v1/move/resources/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
