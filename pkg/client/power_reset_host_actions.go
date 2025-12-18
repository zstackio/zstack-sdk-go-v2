// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PowerResetHost operates on PowerResetHost
func (cli *ZSClient) PowerResetHost(uuid string, params param.PowerResetHostParam) (*view.PowerResetHostEventView, error) {
	resp := view.PowerResetHostEventView{}
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
