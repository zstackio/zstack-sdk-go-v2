// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAccount updates Account
func (cli *ZSClient) UpdateAccount(uuid string, params param.UpdateAccountParam) (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.PutWithRespKey("v1/accounts", uuid, "", map[string]interface{}{
		"updateAccount": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAccount deletes Account
func (cli *ZSClient) DeleteAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts", uuid, string(deleteMode))
}
// CreateAccount creates Account
func (cli *ZSClient) CreateAccount(params param.CreateAccountParam) (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.Post("v1/accounts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAccount queries Account list
func (cli *ZSClient) QueryAccount(params *param.QueryParam) ([]view.AccountInventoryView, error) {
	var resp []view.AccountInventoryView
	return resp, cli.List("v1/accounts", params, &resp)
}

func (cli *ZSClient) GetAccount(uuid string) (*view.AccountInventoryView, error) {
	var resp view.AccountInventoryView
	if err := cli.Get("v1/accounts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccount Pagination
func (cli *ZSClient) PageAccount(params *param.QueryParam) ([]view.AccountInventoryView, int, error) {
	var accounts []view.AccountInventoryView
	total, err := cli.Page("v1/accounts", params, &accounts)
	return accounts, total, err
}

// LogOut revokes the specified session, or the client's session when uuid is empty.
func (cli *ZSClient) LogOut(uuid string, deleteMode param.DeleteMode) error {
	return cli.LogOutWithParams(uuid, deleteMode, param.LogOutParamDetail{})
}

// LogOutWithParams revokes a session and provides client information to Cloud.
func (cli *ZSClient) LogOutWithParams(sessionUuid string, deleteMode param.DeleteMode, params param.LogOutParamDetail) error {
	if sessionUuid == "" {
		sessionUuid = cli.sessionId
	}
	query := url.Values{}
	for key, value := range params.ClientInfo {
		if strings.TrimSpace(key) == "" || strings.Contains(key, ".") {
			return fmt.Errorf("clientInfo keys must be nonempty and must not contain dots")
		}
		query.Set("clientInfo."+key, value)
	}
	return cli.deleteResourceWithQuery("v1/accounts/sessions", sessionUuid, deleteMode, query)
}

// RefreshCaptcha operates on Captcha
func (cli *ZSClient) RefreshCaptcha() (*view.RefreshCaptchaView, error) {
	var resp view.RefreshCaptchaView
	if err := cli.GetWithRespKey("v1/captcha/refresh", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanInvalidLdapBinding operates on InvalidLdapBinding
func (cli *ZSClient) CleanInvalidLdapBinding() (*view.AccountInventoryView, error) {
	resp := view.AccountInventoryView{}
	if err := cli.PutWithRespKey("v1/ldap/bindings/actions", "", "", map[string]interface{}{
		"cleanInvalidLdapBinding": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckApiPermission operates on ApiPermission
func (cli *ZSClient) CheckApiPermission(params param.CheckApiPermissionParam) (*view.CheckApiPermissionView, error) {
	resp := view.CheckApiPermissionView{}
	if err := cli.PutWithRespKey("v1/accounts/permissions/actions", "", "", map[string]interface{}{
		"checkApiPermission": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnlockIdentity operates on UnlockIdentity
func (cli *ZSClient) UnlockIdentity() (*view.UnlockIdentityView, error) {
	var resp view.UnlockIdentityView
	if err := cli.GetWithRespKey("v1/login/control/unlock", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAccountQuotaUsage gets AccountQuotaUsage by uuid
func (cli *ZSClient) GetAccountQuotaUsage(uuid string) (*view.GetAccountQuotaUsageView, error) {
	var resp view.GetAccountQuotaUsageView
	if err := cli.GetWithRespKey("v1/accounts/quota", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidatePassword operates on Password
func (cli *ZSClient) ValidatePassword(params param.ValidatePasswordParam) (*view.ValidatePasswordView, error) {
	resp := view.ValidatePasswordView{}
	if err := cli.PutWithRespKey("v1/password/verify", "", "", map[string]interface{}{
		"validatePassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceAccount gets ResourceAccount by uuid
func (cli *ZSClient) GetResourceAccount() (*view.StringView, error) {
	var resp view.StringView
	if err := cli.GetWithRespKey("v1/resources/accounts", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckResourcePermission operates on ResourcePermission
func (cli *ZSClient) CheckResourcePermission() (*view.CheckResourcePermissionView, error) {
	var resp view.CheckResourcePermissionView
	if err := cli.GetWithRespKey("v1/accounts/resource/api-permissions", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CalculateAccountSpending operates on AccountSpending
func (cli *ZSClient) CalculateAccountSpending(accountUuid string, params param.CalculateAccountSpendingParam) (*view.CalculateAccountSpendingView, error) {
	resp := view.CalculateAccountSpendingView{}
	if err := cli.PutWithRespKey("v1/billings/accounts", accountUuid, "", map[string]interface{}{
		"calculateAccountSpending": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewSession operates on RenewSession
func (cli *ZSClient) RenewSession(sessionUuid string, params param.RenewSessionParam) (*view.SessionInventoryView, error) {
	resp := view.SessionInventoryView{}
	if err := cli.PutWithRespKey("v1/accounts/sessions", sessionUuid, "", map[string]interface{}{
		"renewSession": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportedIdentityModels gets SupportedIdentityModels by uuid
func (cli *ZSClient) GetSupportedIdentityModels() (*view.GetSupportedIdentityModelsView, error) {
	var resp view.GetSupportedIdentityModelsView
	if err := cli.GetWithRespKey("v1/identity-models", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
