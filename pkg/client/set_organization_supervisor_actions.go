// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetOrganizationSupervisor operates on SetOrganizationSupervisor
func (cli *ZSClient) SetOrganizationSupervisor(uuid string, params param.SetOrganizationSupervisorParam) (*view.SetOrganizationSupervisorEventView, error) {
	resp := view.SetOrganizationSupervisorEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
