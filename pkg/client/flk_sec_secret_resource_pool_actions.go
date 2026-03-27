// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateFlkSecSecretResourcePool updates FlkSecSecretResourcePool
func (cli *ZSClient) UpdateFlkSecSecretResourcePool(ctx context.Context, uuid string, params param.UpdateFlkSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/secret-resource-pools/flkSec", uuid, "", map[string]interface{}{
		"updateFlkSecSecretResourcePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateFlkSecSecretResourcePool creates FlkSecSecretResourcePool
func (cli *ZSClient) CreateFlkSecSecretResourcePool(ctx context.Context, params param.CreateFlkSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.Post(ctx, "v1/secret-resource-pool/flkSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
