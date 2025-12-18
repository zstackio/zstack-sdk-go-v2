// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIAM2VirtualIDGroup updates IAM2VirtualIDGroup
func (cli *ZSClient) UpdateIAM2VirtualIDGroup(uuid string, params param.UpdateIAM2VirtualIDGroupParam) (*view.UpdateIAM2VirtualIDGroupEventView, error) {
	resp := view.UpdateIAM2VirtualIDGroupEventView{}
	if err := cli.Put("v1/iam2/projects/groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
