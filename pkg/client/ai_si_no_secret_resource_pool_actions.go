// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAiSiNoSecretResourcePool creates AiSiNoSecretResourcePool
func (cli *ZSClient) CreateAiSiNoSecretResourcePool(ctx context.Context, params param.CreateAiSiNoSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.Post(ctx, "v1/secret-resource-pool/aisino", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
