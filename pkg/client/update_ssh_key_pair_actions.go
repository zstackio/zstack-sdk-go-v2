// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSshKeyPair updates SshKeyPair
func (cli *ZSClient) UpdateSshKeyPair(uuid string, params param.UpdateSshKeyPairParam) (*view.UpdateSshKeyPairEventView, error) {
	resp := view.UpdateSshKeyPairEventView{}
	if err := cli.Put("v1/ssh-key-pair/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
