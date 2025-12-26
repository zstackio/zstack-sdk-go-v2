// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateFlkSecSecretResourcePool updates FlkSecSecretResourcePool
func (cli *ZSClient) UpdateFlkSecSecretResourcePool(uuid string, params param.UpdateFlkSecSecretResourcePoolParam) (*view.UpdateSecretResourcePoolEventView, error) {
	resp := view.UpdateSecretResourcePoolEventView{}
	if err := cli.Put("v1/secret-resource-pools/flkSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
