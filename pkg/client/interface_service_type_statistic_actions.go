// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetInterfaceServiceTypeStatistic 获取InterfaceServiceTypeStatistic详情
func (cli *ZSClient) GetInterfaceServiceTypeStatistic(uuid string) (*view.GetInterfaceServiceTypeStatisticView, error) {
	var resp view.GetInterfaceServiceTypeStatisticView
	if err := cli.Get("v1/hosts/hosts-network-interfaces/service-type-statistic", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

