// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateResourcePrice creates ResourcePrice
func (cli *ZSClient) CreateResourcePrice(params param.CreateResourcePriceParam) (*view.CreateResourcePriceEventView, error) {
	resp := view.CreateResourcePriceEventView{}
	if err := cli.Post("v1/billings/prices", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
