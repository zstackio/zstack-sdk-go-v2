// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVmNicSecurityPolicy changes VmNicSecurityPolicy
func (cli *ZSClient) ChangeVmNicSecurityPolicy(uuid string, params param.ChangeVmNicSecurityPolicyParam) (*view.ChangeVmNicSecurityPolicyEventView, error) {
	resp := view.ChangeVmNicSecurityPolicyEventView{}
	if err := cli.Put("v1/security-groups/nics/{vmNicUuid}/security-policy/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
