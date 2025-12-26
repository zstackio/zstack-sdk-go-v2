// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CalculateAccountSpending operates on CalculateAccountSpending
func (cli *ZSClient) CalculateAccountSpending(uuid string, params param.CalculateAccountSpendingParam) (*view.CalculateAccountSpendingView, error) {
	resp := view.CalculateAccountSpendingView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
