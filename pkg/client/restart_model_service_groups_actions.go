// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RestartModelServiceGroups operates on RestartModelServiceGroups
func (cli *ZSClient) RestartModelServiceGroups(uuid string, params param.RestartModelServiceGroupsParam) (*view.RestartModelServiceGroupsEventView, error) {
	resp := view.RestartModelServiceGroupsEventView{}
	if err := cli.Put("v1/model-service-instance-groups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
