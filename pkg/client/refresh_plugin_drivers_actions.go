// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshPluginDrivers operates on RefreshPluginDrivers
func (cli *ZSClient) RefreshPluginDrivers(uuid string, params param.RefreshPluginDriversParam) (*view.RefreshPluginDriversEventView, error) {
	resp := view.RefreshPluginDriversEventView{}
	if err := cli.Put("v1/external/plugins", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
