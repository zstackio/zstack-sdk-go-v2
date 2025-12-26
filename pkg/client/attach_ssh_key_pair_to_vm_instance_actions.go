// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachSshKeyPairToVmInstance operates on SshKeyPairToVmInstance
func (cli *ZSClient) AttachSshKeyPairToVmInstance(params param.AttachSshKeyPairToVmInstanceParam) (*view.AttachSshKeyPairToVmInstanceEventView, error) {
	resp := view.AttachSshKeyPairToVmInstanceEventView{}
	if err := cli.Post("v1/ssh-key-pair/{sshKeyPairUuid}/vm-instance/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
