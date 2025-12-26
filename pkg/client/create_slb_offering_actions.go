// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSlbOffering creates SlbOffering
func (cli *ZSClient) CreateSlbOffering(params param.CreateSlbOfferingParam) (*view.CreateInstanceOfferingEventView, error) {
	resp := view.CreateInstanceOfferingEventView{}
	if err := cli.Post("v1/instance-offerings/slb", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
