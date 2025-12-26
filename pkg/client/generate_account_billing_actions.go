// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GenerateAccountBilling operates on GenerateAccountBilling
func (cli *ZSClient) GenerateAccountBilling(uuid string, params param.GenerateAccountBillingParam) (*view.GenerateAccountBillingEventView, error) {
	resp := view.GenerateAccountBillingEventView{}
	if err := cli.Put("v1/billings/accounts/{accountUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
