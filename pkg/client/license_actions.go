// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteLicense deletes License
func (cli *ZSClient) DeleteLicense(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/licenses/mn", uuid, string(deleteMode))
}
// UpdateLicense updates License
func (cli *ZSClient) UpdateLicense(managementNodeUuid string, params param.UpdateLicenseParam) (*view.LicenseInventoryView, error) {
	resp := view.LicenseInventoryView{}
	if err := cli.PutWithRespKey("v1/licenses/mn", managementNodeUuid, "", map[string]interface{}{
		"updateLicense": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RegisterLicenseServer operates on LicenseServer
func (cli *ZSClient) RegisterLicenseServer(params param.RegisterLicenseServerParam) (*view.RegisterLicenseServerEventView, error) {
	resp := view.RegisterLicenseServerEventView{}
	if err := cli.Post("v1/license-server/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReloadLicense operates on ReloadLicense
func (cli *ZSClient) ReloadLicense(params param.ReloadLicenseParam) (*view.LicenseInventoryView, error) {
	resp := view.LicenseInventoryView{}
	if err := cli.PutWithRespKey("v1/licenses/actions", "", "", map[string]interface{}{
		"reloadLicense": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IsLicenseServer operates on IsLicenseServer
func (cli *ZSClient) IsLicenseServer() (*view.IsLicenseServerView, error) {
	var resp view.IsLicenseServerView
	if err := cli.GetWithRespKey("v1/license-server/is-server", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseAddOns gets LicenseAddOns by uuid
func (cli *ZSClient) GetLicenseAddOns() (*view.GetLicenseAddOnsView, error) {
	var resp view.GetLicenseAddOnsView
	if err := cli.GetWithRespKey("v1/licenses/addons", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DegradeFromLicenseServer operates on DegradeFromLicenseServer
func (cli *ZSClient) DegradeFromLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server", uuid, string(deleteMode))
}

// PushLicenseAddOnsUsage operates on PushLicenseAddOnsUsage
func (cli *ZSClient) PushLicenseAddOnsUsage(params param.PushLicenseAddOnsUsageParam) (*view.PushLicenseAddOnsUsageEventView, error) {
	resp := view.PushLicenseAddOnsUsageEventView{}
	if err := cli.PutWithRespKey("v1/licenses/addons/usage", "", "", map[string]interface{}{
		"pushLicenseAddOnsUsage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseRecords gets LicenseRecords by uuid
func (cli *ZSClient) GetLicenseRecords() (*view.LicenseInventoryView, error) {
	var resp view.LicenseInventoryView
	if err := cli.GetWithRespKey("v1/licenses/records", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnregisterLicenseRequestedApplication operates on LicenseRequestedApplication
func (cli *ZSClient) UnregisterLicenseRequestedApplication(params param.UnregisterLicenseRequestedApplicationParam) (*view.UnregisterLicenseRequestedApplicationEventView, error) {
	resp := view.UnregisterLicenseRequestedApplicationEventView{}
	if err := cli.PutWithRespKey("v1/license/unregister-applications", "", "", map[string]interface{}{
		"unregisterLicenseRequestedApplication": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseCapabilities gets LicenseCapabilities by uuid
func (cli *ZSClient) GetLicenseCapabilities() (*view.GetLicenseCapabilitiesView, error) {
	var resp view.GetLicenseCapabilitiesView
	if err := cli.GetWithRespKey("v1/licenses/capabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseInfo gets LicenseInfo by uuid
func (cli *ZSClient) GetLicenseInfo() (*view.LicenseInventoryView, error) {
	var resp view.GetLicenseInfoView
	if err := cli.GetWithRespKey("v1/licenses", "", "inventory", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// UnregisterLicenseServer operates on LicenseServer
func (cli *ZSClient) UnregisterLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/client", uuid, string(deleteMode))
}

// GetLicenseNodeUsageDetails gets LicenseNodeUsageDetails by uuid
func (cli *ZSClient) GetLicenseNodeUsageDetails() (*view.GetLicenseNodeUsageDetailsView, error) {
	var resp view.GetLicenseNodeUsageDetailsView
	if err := cli.GetWithRespKey("v1/license/node/usage/details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLicenseUKeyStatus gets LicenseUKeyStatus by uuid
func (cli *ZSClient) GetLicenseUKeyStatus() (*view.UKeyInventoryView, error) {
	resp := view.UKeyInventoryView{}
	if err := cli.Post("v1/licenses/actions", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// WithdrawLicenseCapacityApplication operates on WithdrawLicenseCapacityApplication
func (cli *ZSClient) WithdrawLicenseCapacityApplication(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/license-server/capacity-application", uuid, string(deleteMode))
}

// VerifyLicenseServer operates on VerifyLicenseServer
func (cli *ZSClient) VerifyLicenseServer(params param.VerifyLicenseServerParam) (*view.VerifyLicenseServerEventView, error) {
	resp := view.VerifyLicenseServerEventView{}
	if err := cli.Post("v1/license-server/register-verify", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RegisterLicenseRequestedApplication operates on LicenseRequestedApplication
func (cli *ZSClient) RegisterLicenseRequestedApplication(params param.RegisterLicenseRequestedApplicationParam) (*view.RegisterLicenseRequestedApplicationEventView, error) {
	resp := view.RegisterLicenseRequestedApplicationEventView{}
	if err := cli.Post("v1/licenses/applications", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
