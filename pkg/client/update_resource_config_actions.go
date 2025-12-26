// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateResourceConfig updates ResourceConfig
func (cli *ZSClient) UpdateResourceConfig(uuid string, params param.UpdateResourceConfigParam) (*view.UpdateResourceConfigEventView, error) {
	resp := view.UpdateResourceConfigEventView{}
	if err := cli.Put("v1/resource-configurations/{category}/{name}/{resourceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
