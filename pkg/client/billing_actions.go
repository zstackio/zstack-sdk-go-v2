// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteBilling needs filters to identify the billing records to delete.
//
// Deprecated: use DeleteBillingWithParams with account or time filters.
func (cli *ZSClient) DeleteBilling(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteBilling requires account or time filters; use DeleteBillingWithParams")
}

// DeleteBillingWithParams deletes billing records matching the supplied filters.
func (cli *ZSClient) DeleteBillingWithParams(deleteMode param.DeleteMode, params param.DeleteBillingParamDetail) error {
	query := url.Values{}
	if params.AccountUuid != nil {
		if strings.TrimSpace(*params.AccountUuid) == "" {
			return fmt.Errorf("accountUuid must not be empty")
		}
		query.Set("accountUuid", *params.AccountUuid)
	}
	if params.StartTime != nil {
		query.Set("startTime", strconv.FormatInt(*params.StartTime, 10))
	}
	if params.EndTime != nil {
		query.Set("endTime", strconv.FormatInt(*params.EndTime, 10))
	}
	if len(query) == 0 {
		return fmt.Errorf("at least one billing filter is required; use DeleteAllBilling to delete every record")
	}
	return cli.deleteActionWithQuery("v1/billings/billings", "", "", deleteModeOrDefault(deleteMode, params.DeleteMode), query)
}

// DeleteAllBilling deletes every billing record.
func (cli *ZSClient) DeleteAllBilling(deleteMode param.DeleteMode) error {
	return cli.deleteActionWithQuery("v1/billings/billings", "", "", deleteMode, url.Values{})
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
