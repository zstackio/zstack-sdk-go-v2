// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachPolicyToRole operates on PolicyToRole
func (cli *ZSClient) AttachPolicyToRole(params param.AttachPolicyToRoleParam) (*view.AttachPolicyToRoleEventView, error) {
	resp := view.AttachPolicyToRoleEventView{}
	if err := cli.Post("v1/identities/policies/{policyUuid}/roles/{roleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
