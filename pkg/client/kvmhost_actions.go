// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateKVMHost 更新KVMHost
func (cli *ZSClient) UpdateKVMHost(uuid string, params param.UpdateKVMHostParam) (*view.UpdateHostEventView, error) {
	resp := view.UpdateHostEventView{}
	if err := cli.Put("v1/hosts/kvm/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

