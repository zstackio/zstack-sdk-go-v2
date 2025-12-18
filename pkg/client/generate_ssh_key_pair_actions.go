// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GenerateSshKeyPair 操作GenerateSshKeyPair
func (cli *ZSClient) GenerateSshKeyPair(params param.GenerateSshKeyPairParam) (*view.GenerateSshKeyPairView, error) {
	resp := view.GenerateSshKeyPairView{}
	if err := cli.Post("v1/ssh-key-pair/generate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

