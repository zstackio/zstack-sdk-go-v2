// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateQuota updates Quota
func (cli *ZSClient) UpdateQuota(uuid string, params param.UpdateQuotaParam) (*view.UpdateQuotaEventView, error) {
	resp := view.UpdateQuotaEventView{}
	if err := cli.Put("v1/accounts/quotas/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
