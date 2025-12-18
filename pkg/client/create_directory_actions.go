// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateDirectory creates Directory
func (cli *ZSClient) CreateDirectory(params param.CreateDirectoryParam) (*view.CreateDirectoryEventView, error) {
	resp := view.CreateDirectoryEventView{}
	if err := cli.Post("v1/create/directory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
