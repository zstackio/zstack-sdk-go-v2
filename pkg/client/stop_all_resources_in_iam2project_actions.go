// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StopAllResourcesInIAM2Project stops AllResourcesInIAM2Project
func (cli *ZSClient) StopAllResourcesInIAM2Project(uuid string, params param.StopAllResourcesInIAM2ProjectParam) (*view.StopAllResourcesInIAM2ProjectEventView, error) {
	resp := view.StopAllResourcesInIAM2ProjectEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
