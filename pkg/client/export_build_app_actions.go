// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExportBuildApp 操作ExportBuildApp
func (cli *ZSClient) ExportBuildApp(uuid string, params param.ExportBuildAppParam) (*view.ExportBuildAppEventView, error) {
	resp := view.ExportBuildAppEventView{}
	if err := cli.Put("v1/appcenter/buildapp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

