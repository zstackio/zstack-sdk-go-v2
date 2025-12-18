// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SecurityMachineEncrypt 操作SecurityMachineEncrypt
func (cli *ZSClient) SecurityMachineEncrypt(params param.SecurityMachineEncryptParam) (*view.SecurityMachineEncryptEventView, error) {
	resp := view.SecurityMachineEncryptEventView{}
	if err := cli.Post("v1/security-machine/encrypt/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

