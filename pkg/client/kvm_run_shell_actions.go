// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// KvmRunShell operates on KvmRunShell
func (cli *ZSClient) KvmRunShell(uuid string, params param.KvmRunShellParam) (*view.KvmRunShellEventView, error) {
	resp := view.KvmRunShellEventView{}
	if err := cli.Put("v1/hosts/kvm/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
