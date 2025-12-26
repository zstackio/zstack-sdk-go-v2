// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateCdpPolicy updates CdpPolicy
func (cli *ZSClient) UpdateCdpPolicy(uuid string, params param.UpdateCdpPolicyParam) (*view.UpdateCdpPolicyEventView, error) {
	resp := view.UpdateCdpPolicyEventView{}
	if err := cli.Put("v1/cdp-backup-storage/policy/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
