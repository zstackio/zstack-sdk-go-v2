// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySshKeyPair queries SshKeyPair list
func (cli *ZSClient) QuerySshKeyPair(params *param.QueryParam) ([]view.SshKeyPairInventoryView, error) {
	var resp []view.SshKeyPairInventoryView
	return resp, cli.List("v1/ssh-key-pair", params, &resp)
}
