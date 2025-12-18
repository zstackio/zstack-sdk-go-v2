// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// StopBaremetalInstance stops BaremetalInstance
func (cli *ZSClient) StopBaremetalInstance(uuid string, params param.StopBaremetalInstanceParam) (*view.StopBaremetalInstanceEventView, error) {
	resp := view.StopBaremetalInstanceEventView{}
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
