// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SecurityMachineEncrypt operates on SecurityMachineEncrypt
func (cli *ZSClient) SecurityMachineEncrypt(params param.SecurityMachineEncryptParam) (*view.SecurityMachineEncryptEventView, error) {
	resp := view.SecurityMachineEncryptEventView{}
	if err := cli.Post("v1/security-machine/encrypt/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
