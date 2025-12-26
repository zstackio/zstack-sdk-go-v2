// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBuildApp updates BuildApp
func (cli *ZSClient) UpdateBuildApp(uuid string, params param.UpdateBuildAppParam) (*view.UpdateBuildAppEventView, error) {
	resp := view.UpdateBuildAppEventView{}
	if err := cli.Put("v1/appcenter/buildapp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
