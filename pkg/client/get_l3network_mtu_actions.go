// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetL3NetworkMtu gets L3NetworkMtu by uuid
func (cli *ZSClient) GetL3NetworkMtu(uuid string) (*view.GetL3NetworkMtuView, error) {
	var resp view.GetL3NetworkMtuView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/mtu", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
