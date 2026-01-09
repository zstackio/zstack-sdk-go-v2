// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryQuota queries Quota list
func (cli *ZSClient) QueryQuota(params *param.QueryParam) ([]view.QuotaInventoryView, error) {
	var resp []view.QuotaInventoryView
	return resp, cli.List("v1/accounts/quotas", params, &resp)
}

func (cli *ZSClient) GetQuota(uuid string) (*view.QuotaInventoryView, error) {
	var resp view.QuotaInventoryView
	if err := cli.Get("v1/accounts/quotas", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateQuota updates Quota
func (cli *ZSClient) UpdateQuota(uuid string, params param.UpdateQuotaParam) (*view.QuotaInventoryView, error) {
	var resp view.UpdateQuotaEventView
	if err := cli.Put("v1/accounts/quotas/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
