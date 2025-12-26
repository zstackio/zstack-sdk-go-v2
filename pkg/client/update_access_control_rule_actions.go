// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAccessControlRule updates AccessControlRule
func (cli *ZSClient) UpdateAccessControlRule(uuid string, params param.UpdateAccessControlRuleParam) (*view.UpdateAccessControlRuleEventView, error) {
	resp := view.UpdateAccessControlRuleEventView{}
	if err := cli.Put("v1/login-control/access-control/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
