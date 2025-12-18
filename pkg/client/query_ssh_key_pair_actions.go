// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySshKeyPair queries SshKeyPair list
func (cli *ZSClient) QuerySshKeyPair(params param.QueryParam) ([]view.SshKeyPairInventoryView, error) {
	var resp []view.SshKeyPairInventoryView
	return resp, cli.List("v1/ssh-key-pair", &params, &resp)
}
