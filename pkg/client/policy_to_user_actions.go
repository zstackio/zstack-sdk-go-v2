// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachPolicyToUser 操作PolicyToUser
func (cli *ZSClient) AttachPolicyToUser(params param.AttachPolicyToUserParam) (*view.AttachPolicyToUserEventView, error) {
	resp := view.AttachPolicyToUserEventView{}
	if err := cli.Post("v1/accounts/users/{userUuid}/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

