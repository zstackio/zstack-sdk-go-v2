// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RebootBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) RebootBaremetalInstance(uuid string, params param.RebootBaremetalInstanceParam) (*view.RebootBaremetalInstanceEventView, error) {
	resp := view.RebootBaremetalInstanceEventView{}
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
