// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeIAM2VirtualIDGroupState changes IAM2VirtualIDGroupState
func (cli *ZSClient) ChangeIAM2VirtualIDGroupState(uuid string, params param.ChangeIAM2VirtualIDGroupStateParam) (*view.ChangeIAM2VirtualIDGroupStateEventView, error) {
	resp := view.ChangeIAM2VirtualIDGroupStateEventView{}
	if err := cli.Put("v1/iam2/projects/groups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
