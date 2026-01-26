// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSanSecSecretResourcePool creates SanSecSecretResourcePool
func (cli *ZSClient) CreateSanSecSecretResourcePool(params param.CreateSanSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.Post("v1/secret-resource-pool/sanSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSanSecSecretResourcePool updates SanSecSecretResourcePool
func (cli *ZSClient) UpdateSanSecSecretResourcePool(uuid string, params param.UpdateSanSecSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.PutWithRespKey("v1/secret-resource-pools/sanSec", uuid, "", map[string]interface{}{
		"updateSanSecSecretResourcePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
