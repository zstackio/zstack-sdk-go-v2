// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CalculateResourceSpending operates on CalculateResourceSpending
func (cli *ZSClient) CalculateResourceSpending(uuid string, params param.CalculateResourceSpendingParam) (*view.CalculateResourceSpendingView, error) {
	resp := view.CalculateResourceSpendingView{}
	if err := cli.Put("v1/billings/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
