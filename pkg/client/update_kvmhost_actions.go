// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateKVMHost updates KVMHost
func (cli *ZSClient) UpdateKVMHost(uuid string, params param.UpdateKVMHostParam) (*view.UpdateHostEventView, error) {
	resp := view.UpdateHostEventView{}
	if err := cli.Put("v1/hosts/kvm/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
