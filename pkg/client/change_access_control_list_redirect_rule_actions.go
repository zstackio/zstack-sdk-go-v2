// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeAccessControlListRedirectRule changes AccessControlListRedirectRule
func (cli *ZSClient) ChangeAccessControlListRedirectRule(uuid string, params param.ChangeAccessControlListRedirectRuleParam) (*view.ChangeAccessControlListRedirectRuleEventView, error) {
	resp := view.ChangeAccessControlListRedirectRuleEventView{}
	if err := cli.Put("v1/access-control-lists/redirectRules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
