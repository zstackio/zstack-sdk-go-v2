// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateDirectory updates Directory
func (cli *ZSClient) UpdateDirectory(uuid string, params param.UpdateDirectoryParam) (*view.UpdateDirectoryEventView, error) {
	resp := view.UpdateDirectoryEventView{}
	if err := cli.Put("v1/update/directory", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
