// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSshKeyPair updates SshKeyPair
func (cli *ZSClient) UpdateSshKeyPair(uuid string, params param.UpdateSshKeyPairParam) (*view.UpdateSshKeyPairEventView, error) {
	resp := view.UpdateSshKeyPairEventView{}
	if err := cli.Put("v1/ssh-key-pair/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
