// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryQuota queries Quota list
func (cli *ZSClient) QueryQuota(ctx context.Context, params *param.QueryParam) ([]view.QuotaInventoryView, error) {
	var resp []view.QuotaInventoryView
	return resp, cli.List(ctx, "v1/accounts/quotas", params, &resp)
}

func (cli *ZSClient) GetQuota(ctx context.Context, uuid string) (*view.QuotaInventoryView, error) {
	var resp view.QuotaInventoryView
	if err := cli.Get(ctx, "v1/accounts/quotas", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageQuota Pagination
func (cli *ZSClient) PageQuota(ctx context.Context, params *param.QueryParam) ([]view.QuotaInventoryView, int, error) {
	var quotas []view.QuotaInventoryView
	total, err := cli.Page(ctx, "v1/accounts/quotas", params, &quotas)
	return quotas, total, err
}
// UpdateQuota updates Quota
func (cli *ZSClient) UpdateQuota(ctx context.Context, params param.UpdateQuotaParam) (*view.QuotaInventoryView, error) {
	resp := view.QuotaInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/accounts/quotas/actions", "", "", map[string]interface{}{
		"updateQuota": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
