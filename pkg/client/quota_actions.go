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

// PageQuota Pagination
func (cli *ZSClient) PageQuota(params *param.QueryParam) ([]view.QuotaInventoryView, int, error) {
	var quotas []view.QuotaInventoryView
	total, err := cli.Page("v1/accounts/quotas", params, &quotas)
	return quotas, total, err
}
// UpdateQuota updates Quota
func (cli *ZSClient) UpdateQuota(uuid string, params param.UpdateQuotaParam) (*view.QuotaInventoryView, error) {
	resp := view.QuotaInventoryView{}
	if err := cli.Put("v1/accounts/quotas/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
