// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetAccountQuotaUsage gets AccountQuotaUsage by uuid
func (cli *ZSClient) GetAccountQuotaUsage(uuid string) (*view.GetAccountQuotaUsageView, error) {
	var resp view.GetAccountQuotaUsageView
	if err := cli.Get("v1/accounts/quota/{uuid}/usages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
