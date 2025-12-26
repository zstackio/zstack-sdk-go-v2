// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeIAM2ProjectState changes IAM2ProjectState
func (cli *ZSClient) ChangeIAM2ProjectState(uuid string, params param.ChangeIAM2ProjectStateParam) (*view.ChangeIAM2ProjectStateEventView, error) {
	resp := view.ChangeIAM2ProjectStateEventView{}
	if err := cli.Put("v1/iam2/projects/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
