// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RestartModelServiceGroups 操作RestartModelServiceGroups
func (cli *ZSClient) RestartModelServiceGroups(uuid string, params param.RestartModelServiceGroupsParam) (*view.RestartModelServiceGroupsEventView, error) {
	resp := view.RestartModelServiceGroupsEventView{}
	if err := cli.Put("v1/model-service-instance-groups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

