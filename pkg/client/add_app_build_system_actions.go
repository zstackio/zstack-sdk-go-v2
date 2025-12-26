// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAppBuildSystem adds AppBuildSystem
func (cli *ZSClient) AddAppBuildSystem(params param.AddAppBuildSystemParam) (*view.AddAppBuildSystemEventView, error) {
	resp := view.AddAppBuildSystemEventView{}
	if err := cli.Post("v1/appcenter/buildsystem", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
