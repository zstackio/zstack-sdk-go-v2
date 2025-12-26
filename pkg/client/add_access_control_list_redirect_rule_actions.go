// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAccessControlListRedirectRule adds AccessControlListRedirectRule
func (cli *ZSClient) AddAccessControlListRedirectRule(params param.AddAccessControlListRedirectRuleParam) (*view.AddAccessControlListEntryEventView, error) {
	resp := view.AddAccessControlListEntryEventView{}
	if err := cli.Post("v1/access-control-lists/{aclUuid}/redirectRules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
