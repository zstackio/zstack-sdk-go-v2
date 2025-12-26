// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GenerateSshKeyPair operates on GenerateSshKeyPair
func (cli *ZSClient) GenerateSshKeyPair(params param.GenerateSshKeyPairParam) (*view.GenerateSshKeyPairView, error) {
	resp := view.GenerateSshKeyPairView{}
	if err := cli.Post("v1/ssh-key-pair/generate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
