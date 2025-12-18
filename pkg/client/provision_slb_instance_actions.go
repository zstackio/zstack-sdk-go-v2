// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ProvisionSlbInstance 操作ProvisionSlbInstance
func (cli *ZSClient) ProvisionSlbInstance(uuid string, params param.ProvisionSlbInstanceParam) (*view.ProvisionSlbGroupInstanceEventView, error) {
	resp := view.ProvisionSlbGroupInstanceEventView{}
	if err := cli.Put("v1/load-balancers/slb/instances/{uuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

