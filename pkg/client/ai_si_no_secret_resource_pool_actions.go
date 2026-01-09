// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateAiSiNoSecretResourcePool creates AiSiNoSecretResourcePool
func (cli *ZSClient) CreateAiSiNoSecretResourcePool(params param.CreateAiSiNoSecretResourcePoolParam) (*view.SecretResourcePoolInventoryView, error) {
	var resp view.CreateSecretResourcePoolEventView
	if err := cli.Post("v1/secret-resource-pool/aisino", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
