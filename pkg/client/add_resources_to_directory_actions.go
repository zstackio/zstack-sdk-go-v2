// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddResourcesToDirectory 操作AddResourcesToDirectory
func (cli *ZSClient) AddResourcesToDirectory(params param.AddResourcesToDirectoryParam) (*view.AddResourcesToDirectoryEventView, error) {
	resp := view.AddResourcesToDirectoryEventView{}
	if err := cli.Post("v1/add/resources/directory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

