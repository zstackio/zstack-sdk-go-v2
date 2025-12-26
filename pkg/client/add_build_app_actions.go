// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddBuildApp adds BuildApp
func (cli *ZSClient) AddBuildApp(params param.AddBuildAppParam) (*view.AddBuildAppEventView, error) {
	resp := view.AddBuildAppEventView{}
	if err := cli.Post("v1/appcenter/buildapp/add", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
