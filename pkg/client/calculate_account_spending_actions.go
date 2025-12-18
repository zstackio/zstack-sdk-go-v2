// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CalculateAccountSpending 操作CalculateAccountSpending
func (cli *ZSClient) CalculateAccountSpending(uuid string, params param.CalculateAccountSpendingParam) (*view.CalculateAccountSpendingView, error) {
	resp := view.CalculateAccountSpendingView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

