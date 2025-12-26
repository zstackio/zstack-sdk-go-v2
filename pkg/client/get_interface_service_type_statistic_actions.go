// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetInterfaceServiceTypeStatistic gets InterfaceServiceTypeStatistic by uuid
func (cli *ZSClient) GetInterfaceServiceTypeStatistic(uuid string) (*view.GetInterfaceServiceTypeStatisticView, error) {
	var resp view.GetInterfaceServiceTypeStatisticView
	if err := cli.Get("v1/hosts/hosts-network-interfaces/service-type-statistic", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
