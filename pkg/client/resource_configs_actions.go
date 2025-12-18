// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateResourceConfigs 更新ResourceConfigs
func (cli *ZSClient) UpdateResourceConfigs(uuid string, params param.UpdateResourceConfigsParam) (*view.UpdateResourceConfigsEventView, error) {
	resp := view.UpdateResourceConfigsEventView{}
	if err := cli.Put("v1/resource-configurations/{resourceUuid}/resource-configs/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

