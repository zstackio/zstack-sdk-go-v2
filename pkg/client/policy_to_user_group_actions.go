// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachPolicyToUserGroup 操作PolicyToUserGroup
func (cli *ZSClient) AttachPolicyToUserGroup(params param.AttachPolicyToUserGroupParam) (*view.AttachPolicyToUserGroupEventView, error) {
	resp := view.AttachPolicyToUserGroupEventView{}
	if err := cli.Post("v1/accounts/groups/{groupUuid}/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

