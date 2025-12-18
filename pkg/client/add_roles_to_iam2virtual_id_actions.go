// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddRolesToIAM2VirtualID 操作AddRolesToIAM2VirtualID
func (cli *ZSClient) AddRolesToIAM2VirtualID(params param.AddRolesToIAM2VirtualIDParam) (*view.AddRolesToIAM2VirtualIDEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDEventView{}
	if err := cli.Post("v1/iam2/projects/virtual-ids/{virtualIDUuid}/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

