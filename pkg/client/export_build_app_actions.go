// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ExportBuildApp operates on ExportBuildApp
func (cli *ZSClient) ExportBuildApp(uuid string, params param.ExportBuildAppParam) (*view.ExportBuildAppEventView, error) {
	resp := view.ExportBuildAppEventView{}
	if err := cli.Put("v1/appcenter/buildapp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
