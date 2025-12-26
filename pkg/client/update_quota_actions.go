// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateQuota updates Quota
func (cli *ZSClient) UpdateQuota(uuid string, params param.UpdateQuotaParam) (*view.UpdateQuotaEventView, error) {
	resp := view.UpdateQuotaEventView{}
	if err := cli.Put("v1/accounts/quotas/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
