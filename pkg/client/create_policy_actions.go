// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreatePolicy creates Policy
func (cli *ZSClient) CreatePolicy(params param.CreatePolicyParam) (*view.CreatePolicyEventView, error) {
	resp := view.CreatePolicyEventView{}
	if err := cli.Post("v1/accounts/policies", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
