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

// DeleteLicense requires a license selector in addition to the management node.
//
// Deprecated: use DeleteLicenseWithParams to specify a license UUID or module.
func (cli *ZSClient) DeleteLicense(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteLicense requires a license UUID or module; use DeleteLicenseWithParams")
}

// DeleteLicenseWithParams deletes a license by UUID, or by module when the UUID is absent.
func (cli *ZSClient) DeleteLicenseWithParams(managementNodeUuid string, deleteMode param.DeleteMode, params param.DeleteLicenseParamDetail) error {
	query := url.Values{}
	required := "uuid"
	if params.Uuid != "" {
		query.Set("uuid", params.Uuid)
	} else {
		required = "module"
	}
	if params.Module != nil {
		query.Set("module", *params.Module)
	}
	return cli.deleteActionWithQuery("v1/licenses/mn", managementNodeUuid, "actions", deleteMode, query, required)
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

// UnregisterLicenseServer unregisters a client node, or its server when uuid is empty.
func (cli *ZSClient) UnregisterLicenseServer(uuid string, deleteMode param.DeleteMode) error {
	params := param.UnregisterLicenseServerParamDetail{}
	if uuid != "" {
		params.ClientAuthorizedNodeUuid = &uuid
	}
	return cli.UnregisterLicenseServerWithParams(deleteMode, params)
}

// UnregisterLicenseServerWithParams unregisters a node or the client's license server.
func (cli *ZSClient) UnregisterLicenseServerWithParams(deleteMode param.DeleteMode, params param.UnregisterLicenseServerParamDetail) error {
	query := url.Values{}
	if params.ClientAuthorizedNodeUuid != nil {
		if strings.TrimSpace(*params.ClientAuthorizedNodeUuid) == "" {
			return fmt.Errorf("clientAuthorizedNodeUuid must not be empty when supplied")
		}
		query.Set("clientAuthorizedNodeUuid", *params.ClientAuthorizedNodeUuid)
	}
	return cli.deleteActionWithQuery("v1/license-server/client", "", "", deleteMode, query)
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

// WithdrawLicenseCapacityApplication requires additional arguments.
//
// Deprecated: use WithdrawLicenseCapacityApplicationWithParams to supply the required fields.
func (cli *ZSClient) WithdrawLicenseCapacityApplication(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("WithdrawLicenseCapacityApplication requires additional arguments; use WithdrawLicenseCapacityApplicationWithParams")
}

// WithdrawLicenseCapacityApplicationWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) WithdrawLicenseCapacityApplicationWithParams(deleteMode param.DeleteMode, params param.WithdrawLicenseCapacityApplicationParamDetail) error {
	query := url.Values{
		"resourceUuidList":         params.ResourceUuidList,
		"clientAuthorizedNodeUuid": {params.ClientAuthorizedNodeUuid},
		"licenseType":              {params.LicenseType},
	}
	return cli.deleteActionWithQuery("v1/license-server/capacity-application", "", "", deleteMode, query, "resourceUuidList", "clientAuthorizedNodeUuid", "licenseType")
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
