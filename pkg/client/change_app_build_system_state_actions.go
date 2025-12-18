// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeAppBuildSystemState changes AppBuildSystemState
func (cli *ZSClient) ChangeAppBuildSystemState(uuid string, params param.ChangeAppBuildSystemStateParam) (*view.ChangeAppBuildSystemStateEventView, error) {
	resp := view.ChangeAppBuildSystemStateEventView{}
	if err := cli.Put("v1/appcenter/buildsystem/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
