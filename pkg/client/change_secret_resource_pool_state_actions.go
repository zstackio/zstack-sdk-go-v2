// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSecretResourcePoolState changes SecretResourcePoolState
func (cli *ZSClient) ChangeSecretResourcePoolState(uuid string, params param.ChangeSecretResourcePoolStateParam) (*view.ChangeSecretResourcePoolStateEventView, error) {
	resp := view.ChangeSecretResourcePoolStateEventView{}
	if err := cli.Put("v1/secret-resource-pools/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
