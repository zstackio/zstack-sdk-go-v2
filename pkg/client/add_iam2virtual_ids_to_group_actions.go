// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIAM2VirtualIDsToGroup adds IAM2VirtualIDsToGroup
func (cli *ZSClient) AddIAM2VirtualIDsToGroup(params param.AddIAM2VirtualIDsToGroupParam) (*view.AddIAM2VirtualIDToGroupEventView, error) {
	resp := view.AddIAM2VirtualIDToGroupEventView{}
	if err := cli.Post("v1/iam2/projects/groups/{groupUuid}/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
