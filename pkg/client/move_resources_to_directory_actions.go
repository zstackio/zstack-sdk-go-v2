// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// MoveResourcesToDirectory 操作MoveResourcesToDirectory
func (cli *ZSClient) MoveResourcesToDirectory(uuid string, params param.MoveResourcesToDirectoryParam) (*view.MoveResourcesToDirectoryEventView, error) {
	resp := view.MoveResourcesToDirectoryEventView{}
	if err := cli.Put("v1/move/resources/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

