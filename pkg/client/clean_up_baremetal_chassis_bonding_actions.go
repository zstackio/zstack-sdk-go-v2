// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanUpBaremetalChassisBonding 操作CleanUpBaremetalChassisBonding
func (cli *ZSClient) CleanUpBaremetalChassisBonding(uuid string, params param.CleanUpBaremetalChassisBondingParam) (*view.CleanUpBaremetalChassisBondingEventView, error) {
	resp := view.CleanUpBaremetalChassisBondingEventView{}
	if err := cli.Put("v1/baremetal/chassis/{chassisUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

