// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectAppBuildSystem operates on ReconnectAppBuildSystem
func (cli *ZSClient) ReconnectAppBuildSystem(uuid string, params param.ReconnectAppBuildSystemParam) (*view.ReconnectAppBuildSystemEventView, error) {
	resp := view.ReconnectAppBuildSystemEventView{}
	if err := cli.Put("v1/appcenter/buildsystem/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
