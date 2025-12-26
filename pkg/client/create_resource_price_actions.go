// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateResourcePrice creates ResourcePrice
func (cli *ZSClient) CreateResourcePrice(params param.CreateResourcePriceParam) (*view.CreateResourcePriceEventView, error) {
	resp := view.CreateResourcePriceEventView{}
	if err := cli.Post("v1/billings/prices", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
