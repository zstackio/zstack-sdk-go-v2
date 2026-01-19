// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateKoAlSecretResourcePool updates KoAlSecretResourcePool
func (cli *ZSClient) UpdateKoAlSecretResourcePool(uuid string, params param.UpdateKoAlSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.Put("v1/secret-resource-pools/koal", uuid, map[string]interface{}{
		"updateKoAlSecretResourcePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateKoAlSecretResourcePool creates KoAlSecretResourcePool
func (cli *ZSClient) CreateKoAlSecretResourcePool(params param.CreateKoAlSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	resp := view.SecretResourcePoolInventoryView{}
	if err := cli.Post("v1/secret-resource-pool/koal", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
