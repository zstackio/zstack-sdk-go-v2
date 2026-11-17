// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteBilling deletes Billing
func (cli *ZSClient) DeleteBilling(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/billings", uuid, string(deleteMode))
}

// GenerateAccountBilling operates on AccountBilling
func (cli *ZSClient) GenerateAccountBilling(accountUuid string, params param.GenerateAccountBillingParam) (*view.GenerateAccountBillingEventView, error) {
	resp := view.GenerateAccountBillingEventView{}
	if err := cli.PutWithRespKey("v1/billings/accounts", accountUuid, "", map[string]interface{}{
		"generateAccountBilling": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanupBillingUsage operates on upBillingUsage
func (cli *ZSClient) CleanupBillingUsage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/usage", uuid, string(deleteMode))
}

// QueryAccountBilling queries AccountBilling list
func (cli *ZSClient) QueryAccountBilling(params *param.QueryParam) ([]view.BillingInventoryView, error) {
	var resp []view.BillingInventoryView
	return resp, cli.List("v1/billing/billings", params, &resp)
}

func (cli *ZSClient) GetAccountBilling(uuid string) (*view.BillingInventoryView, error) {
	var resp view.BillingInventoryView
	if err := cli.Get("v1/billing/billings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccountBilling Pagination
func (cli *ZSClient) PageAccountBilling(params *param.QueryParam) ([]view.BillingInventoryView, int, error) {
	var accountBillings []view.BillingInventoryView
	total, err := cli.Page("v1/billing/billings", params, &accountBillings)
	return accountBillings, total, err
}

// CalculateResourceSpending operates on ResourceSpending
func (cli *ZSClient) CalculateResourceSpending(params param.CalculateResourceSpendingParam) (*view.CalculateResourceSpendingView, error) {
	resp := view.CalculateResourceSpendingView{}
	if err := cli.PutWithRespKey("v1/billings/resources/actions", "", "", map[string]interface{}{
		"calculateResourceSpending": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateAccountBillingSpending operates on AccountBillingSpending
func (cli *ZSClient) CalculateAccountBillingSpending(accountUuid string, params param.CalculateAccountBillingSpendingParam) (*view.CalculateAccountBillingSpendingView, error) {
	resp := view.CalculateAccountBillingSpendingView{}
	if err := cli.PutWithRespKey("v1/billings/accounts", accountUuid, "", map[string]interface{}{
		"calculateAccountBillingSpending": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
