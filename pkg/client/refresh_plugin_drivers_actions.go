// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshPluginDrivers 操作RefreshPluginDrivers
func (cli *ZSClient) RefreshPluginDrivers(uuid string, params param.RefreshPluginDriversParam) (*view.RefreshPluginDriversEventView, error) {
	resp := view.RefreshPluginDriversEventView{}
	if err := cli.Put("v1/external/plugins", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

