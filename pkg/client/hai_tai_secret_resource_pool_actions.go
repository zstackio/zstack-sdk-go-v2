// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateHaiTaiSecretResourcePool creates HaiTaiSecretResourcePool
func (cli *ZSClient) CreateHaiTaiSecretResourcePool(ctx context.Context, params param.CreateHaiTaiSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/secret-resource-pool/haitai", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
