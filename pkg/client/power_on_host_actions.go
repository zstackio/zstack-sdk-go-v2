// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PowerOnHost operates on PowerOnHost
func (cli *ZSClient) PowerOnHost(uuid string, params param.PowerOnHostParam) (*view.PowerOnHostEventView, error) {
	resp := view.PowerOnHostEventView{}
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
