// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryQuota queries Quota list
func (cli *ZSClient) QueryQuota(params *param.QueryParam) ([]view.QuotaInventoryView, error) {
	var resp []view.QuotaInventoryView
	return resp, cli.List("v1/accounts/quotas", params, &resp)
}
// UpdateQuota updates Quota
func (cli *ZSClient) UpdateQuota(uuid string, params param.UpdateQuotaParam) (*view.QuotaInventoryView, error) {
	var resp view.UpdateQuotaEventView
	if err := cli.Put("v1/accounts/quotas/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
