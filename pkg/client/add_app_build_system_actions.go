// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAppBuildSystem 操作AddAppBuildSystem
func (cli *ZSClient) AddAppBuildSystem(params param.AddAppBuildSystemParam) (*view.AddAppBuildSystemEventView, error) {
	resp := view.AddAppBuildSystemEventView{}
	if err := cli.Post("v1/appcenter/buildsystem", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

