// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryQuota queries Quota list
func (cli *ZSClient) QueryQuota(params *param.QueryParam) ([]view.QuotaInventoryView, error) {
	var resp []view.QuotaInventoryView
	return resp, cli.List("v1/accounts/quotas", params, &resp)
}
