// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetOrganizationSupervisor 操作SetOrganizationSupervisor
func (cli *ZSClient) SetOrganizationSupervisor(uuid string, params param.SetOrganizationSupervisorParam) (*view.SetOrganizationSupervisorEventView, error) {
	resp := view.SetOrganizationSupervisorEventView{}
	if err := cli.Put("v1/iam2/organizations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

