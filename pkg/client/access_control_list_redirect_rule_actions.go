// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeAccessControlListRedirectRule 操作AccessControlListRedirectRule
func (cli *ZSClient) ChangeAccessControlListRedirectRule(uuid string, params param.ChangeAccessControlListRedirectRuleParam) (*view.ChangeAccessControlListRedirectRuleEventView, error) {
	resp := view.ChangeAccessControlListRedirectRuleEventView{}
	if err := cli.Put("v1/access-control-lists/redirectRules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

