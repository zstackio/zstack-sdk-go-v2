// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateResourceConfigs updates ResourceConfigs
func (cli *ZSClient) UpdateResourceConfigs(uuid string, params param.UpdateResourceConfigsParam) (*view.UpdateResourceConfigsEventView, error) {
	resp := view.UpdateResourceConfigsEventView{}
	if err := cli.Put("v1/resource-configurations/{resourceUuid}/resource-configs/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
