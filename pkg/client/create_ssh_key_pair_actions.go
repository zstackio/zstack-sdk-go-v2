// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSshKeyPair creates SshKeyPair
func (cli *ZSClient) CreateSshKeyPair(params param.CreateSshKeyPairParam) (*view.CreateSshKeyPairEventView, error) {
	resp := view.CreateSshKeyPairEventView{}
	if err := cli.Post("v1/ssh-key-pair", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
