// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// StopAllResourcesInIAM2Project stops AllResourcesInIAM2Project
func (cli *ZSClient) StopAllResourcesInIAM2Project(uuid string, params param.StopAllResourcesInIAM2ProjectParam) (*view.StopAllResourcesInIAM2ProjectEventView, error) {
	resp := view.StopAllResourcesInIAM2ProjectEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
