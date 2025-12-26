// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateModelServiceInstanceGroup updates ModelServiceInstanceGroup
func (cli *ZSClient) UpdateModelServiceInstanceGroup(uuid string, params param.UpdateModelServiceInstanceGroupParam) (*view.UpdateModelServiceInstanceGroupEventView, error) {
	resp := view.UpdateModelServiceInstanceGroupEventView{}
	if err := cli.Put("v1/model-service-instance-groups/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
