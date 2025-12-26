// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateResourceStackFromApp creates ResourceStackFromApp
func (cli *ZSClient) CreateResourceStackFromApp(params param.CreateResourceStackFromAppParam) (*view.CreateResourceStackEventView, error) {
	resp := view.CreateResourceStackEventView{}
	if err := cli.Post("v1/appcenter/app/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
