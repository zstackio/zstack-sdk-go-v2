// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachPolicyToRole 操作PolicyToRole
func (cli *ZSClient) AttachPolicyToRole(params param.AttachPolicyToRoleParam) (*view.AttachPolicyToRoleEventView, error) {
	resp := view.AttachPolicyToRoleEventView{}
	if err := cli.Post("v1/identities/policies/{policyUuid}/roles/{roleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

