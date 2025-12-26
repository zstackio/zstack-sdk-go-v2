// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAppBuildSystem updates AppBuildSystem
func (cli *ZSClient) UpdateAppBuildSystem(uuid string, params param.UpdateAppBuildSystemParam) (*view.UpdateAppBuildSystemEventView, error) {
	resp := view.UpdateAppBuildSystemEventView{}
	if err := cli.Put("v1/appcenter/buildsystem/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
