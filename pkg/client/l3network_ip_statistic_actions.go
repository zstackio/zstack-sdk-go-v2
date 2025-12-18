// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetL3NetworkIpStatistic 获取L3NetworkIpStatistic详情
func (cli *ZSClient) GetL3NetworkIpStatistic(uuid string) (*view.GetL3NetworkIpStatisticView, error) {
	var resp view.GetL3NetworkIpStatisticView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/ip-statistic", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

