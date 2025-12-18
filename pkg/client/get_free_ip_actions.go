// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetFreeIp gets FreeIp by uuid
func (cli *ZSClient) GetFreeIp(uuid string) (*view.GetFreeIpView, error) {
	var resp view.GetFreeIpView
	if err := cli.Get("v1/l3-networks/ip/free", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
