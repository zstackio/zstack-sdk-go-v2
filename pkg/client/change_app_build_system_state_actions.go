// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeAppBuildSystemState changes AppBuildSystemState
func (cli *ZSClient) ChangeAppBuildSystemState(uuid string, params param.ChangeAppBuildSystemStateParam) (*view.ChangeAppBuildSystemStateEventView, error) {
	resp := view.ChangeAppBuildSystemStateEventView{}
	if err := cli.Put("v1/appcenter/buildsystem/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
