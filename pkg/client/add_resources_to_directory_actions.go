// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddResourcesToDirectory adds ResourcesToDirectory
func (cli *ZSClient) AddResourcesToDirectory(params param.AddResourcesToDirectoryParam) (*view.AddResourcesToDirectoryEventView, error) {
	resp := view.AddResourcesToDirectoryEventView{}
	if err := cli.Post("v1/add/resources/directory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
