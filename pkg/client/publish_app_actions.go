// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PublishApp operates on PublishApp
func (cli *ZSClient) PublishApp(params param.PublishAppParam) (*view.PublishAppEventView, error) {
	resp := view.PublishAppEventView{}
	if err := cli.Post("v1/appcenter/app", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
