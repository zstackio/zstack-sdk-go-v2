// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAccessControlListRedirectRule 操作AddAccessControlListRedirectRule
func (cli *ZSClient) AddAccessControlListRedirectRule(params param.AddAccessControlListRedirectRuleParam) (*view.AddAccessControlListEntryEventView, error) {
	resp := view.AddAccessControlListEntryEventView{}
	if err := cli.Post("v1/access-control-lists/{aclUuid}/redirectRules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

