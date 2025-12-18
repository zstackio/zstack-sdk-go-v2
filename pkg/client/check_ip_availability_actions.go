// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckIpAvailability 操作CheckIpAvailability
func (cli *ZSClient) CheckIpAvailability(params param.CheckIpAvailabilityParam) (*view.CheckIpAvailabilityView, error) {
	var resp view.CheckIpAvailabilityView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/ip/{ip}/availability", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

