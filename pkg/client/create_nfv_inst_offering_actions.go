// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateNfvInstOffering creates NfvInstOffering
func (cli *ZSClient) CreateNfvInstOffering(params param.CreateNfvInstOfferingParam) (*view.CreateInstanceOfferingEventView, error) {
	resp := view.CreateInstanceOfferingEventView{}
	if err := cli.Post("v1/instance-offerings/nfvinst", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
