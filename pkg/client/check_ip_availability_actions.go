// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckIpAvailability operates on CheckIpAvailability
func (cli *ZSClient) CheckIpAvailability(params param.CheckIpAvailabilityParam) (*view.CheckIpAvailabilityView, error) {
	var resp view.CheckIpAvailabilityView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/ip/{ip}/availability", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
