// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAttributesToIAM2VirtualIDGroup adds AttributesToIAM2VirtualIDGroup
func (cli *ZSClient) AddAttributesToIAM2VirtualIDGroup(params param.AddAttributesToIAM2VirtualIDGroupParam) (*view.AddAttributesToIAM2VirtualIDGroupEventView, error) {
	resp := view.AddAttributesToIAM2VirtualIDGroupEventView{}
	if err := cli.Post("v1/iam2/projects/groups/{uuid}/attributes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
