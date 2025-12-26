// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateIAM2VirtualIDGroupAttribute updates IAM2VirtualIDGroupAttribute
func (cli *ZSClient) UpdateIAM2VirtualIDGroupAttribute(uuid string, params param.UpdateIAM2VirtualIDGroupAttributeParam) (*view.UpdateIAM2VirtualIDGroupAttributeEventView, error) {
	resp := view.UpdateIAM2VirtualIDGroupAttributeEventView{}
	if err := cli.Put("v1/iam2/projects/groups/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
