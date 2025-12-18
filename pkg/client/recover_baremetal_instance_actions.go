// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RecoverBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) RecoverBaremetalInstance(uuid string, params param.RecoverBaremetalInstanceParam) (*view.RecoverBaremetalInstanceEventView, error) {
	resp := view.RecoverBaremetalInstanceEventView{}
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
