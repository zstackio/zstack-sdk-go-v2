// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckVipPortAvailability operates on CheckVipPortAvailability
func (cli *ZSClient) CheckVipPortAvailability(params param.CheckVipPortAvailabilityParam) (*view.CheckVipPortAvailabilityView, error) {
	var resp view.CheckVipPortAvailabilityView
	if err := cli.Get("v1/vips/{vipUuid}/check-port-availability", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
