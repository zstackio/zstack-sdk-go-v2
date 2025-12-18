// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateL3Network creates L3Network
func (cli *ZSClient) CreateL3Network(params param.CreateL3NetworkParam) (*view.CreateL3NetworkEventView, error) {
	resp := view.CreateL3NetworkEventView{}
	if err := cli.Post("v1/l3-networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
