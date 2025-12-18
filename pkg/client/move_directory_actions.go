// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// MoveDirectory operates on MoveDirectory
func (cli *ZSClient) MoveDirectory(uuid string, params param.MoveDirectoryParam) (*view.MoveDirectoryEventView, error) {
	resp := view.MoveDirectoryEventView{}
	if err := cli.Put("v1/move/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
