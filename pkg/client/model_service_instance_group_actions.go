// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateModelServiceInstanceGroup 更新ModelServiceInstanceGroup
func (cli *ZSClient) UpdateModelServiceInstanceGroup(uuid string, params param.UpdateModelServiceInstanceGroupParam) (*view.UpdateModelServiceInstanceGroupEventView, error) {
	resp := view.UpdateModelServiceInstanceGroupEventView{}
	if err := cli.Put("v1/model-service-instance-groups/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

