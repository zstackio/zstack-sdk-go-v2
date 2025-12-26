// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateCdpPolicy creates CdpPolicy
func (cli *ZSClient) CreateCdpPolicy(params param.CreateCdpPolicyParam) (*view.CreateCdpPolicyEventView, error) {
	resp := view.CreateCdpPolicyEventView{}
	if err := cli.Post("v1/cdp-backup-storage/policy", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
