// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAccessControlRule adds AccessControlRule
func (cli *ZSClient) AddAccessControlRule(params param.AddAccessControlRuleParam) (*view.AddAccessControlRuleEventView, error) {
	resp := view.AddAccessControlRuleEventView{}
	if err := cli.Post("v1/login-control/access-control/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
