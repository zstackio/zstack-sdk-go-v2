// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CalculateAccountBillingSpending 操作CalculateAccountBillingSpending
func (cli *ZSClient) CalculateAccountBillingSpending(uuid string, params param.CalculateAccountBillingSpendingParam) (*view.CalculateAccountBillingSpendingView, error) {
	resp := view.CalculateAccountBillingSpendingView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

