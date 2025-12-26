// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachPolicyToUser operates on PolicyToUser
func (cli *ZSClient) AttachPolicyToUser(params param.AttachPolicyToUserParam) (*view.AttachPolicyToUserEventView, error) {
	resp := view.AttachPolicyToUserEventView{}
	if err := cli.Post("v1/accounts/users/{userUuid}/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
