// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySecretResourcePool queries SecretResourcePool list
func (cli *ZSClient) QuerySecretResourcePool(params *param.QueryParam) ([]view.SecretResourcePoolInventoryView, error) {
	var resp []view.SecretResourcePoolInventoryView
	return resp, cli.List("v1/secret-resource-pools", params, &resp)
}
