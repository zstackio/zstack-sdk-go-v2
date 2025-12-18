// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachSshKeyPairToVmInstance operates on SshKeyPairToVmInstance
func (cli *ZSClient) AttachSshKeyPairToVmInstance(params param.AttachSshKeyPairToVmInstanceParam) (*view.AttachSshKeyPairToVmInstanceEventView, error) {
	resp := view.AttachSshKeyPairToVmInstanceEventView{}
	if err := cli.Post("v1/ssh-key-pair/{sshKeyPairUuid}/vm-instance/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
