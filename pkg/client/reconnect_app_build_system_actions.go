// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectAppBuildSystem 操作ReconnectAppBuildSystem
func (cli *ZSClient) ReconnectAppBuildSystem(uuid string, params param.ReconnectAppBuildSystemParam) (*view.ReconnectAppBuildSystemEventView, error) {
	resp := view.ReconnectAppBuildSystemEventView{}
	if err := cli.Put("v1/appcenter/buildsystem/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

