// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachPolicyToUserGroup operates on PolicyToUserGroup
func (cli *ZSClient) AttachPolicyToUserGroup(params param.AttachPolicyToUserGroupParam) (*view.AttachPolicyToUserGroupEventView, error) {
	resp := view.AttachPolicyToUserGroupEventView{}
	if err := cli.Post("v1/accounts/groups/{groupUuid}/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
