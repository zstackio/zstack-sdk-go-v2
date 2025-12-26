// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateDirectory updates Directory
func (cli *ZSClient) UpdateDirectory(uuid string, params param.UpdateDirectoryParam) (*view.UpdateDirectoryEventView, error) {
	resp := view.UpdateDirectoryEventView{}
	if err := cli.Put("v1/update/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
