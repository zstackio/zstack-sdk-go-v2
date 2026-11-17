// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2Organization queries IAM2Organization list
func (cli *ZSClient) QueryIAM2Organization(params *param.QueryParam) ([]view.IAM2OrganizationInventoryView, error) {
	var resp []view.IAM2OrganizationInventoryView
	return resp, cli.List("v1/iam2/organizations", params, &resp)
}

func (cli *ZSClient) GetIAM2Organization(uuid string) (*view.IAM2OrganizationInventoryView, error) {
	var resp view.IAM2OrganizationInventoryView
	if err := cli.Get("v1/iam2/organizations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2Organization Pagination
func (cli *ZSClient) PageIAM2Organization(params *param.QueryParam) ([]view.IAM2OrganizationInventoryView, int, error) {
	var iAM2Organizations []view.IAM2OrganizationInventoryView
	total, err := cli.Page("v1/iam2/organizations", params, &iAM2Organizations)
	return iAM2Organizations, total, err
}
// DeleteIAM2Organization deletes IAM2Organization
func (cli *ZSClient) DeleteIAM2Organization(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/organizations", uuid, string(deleteMode))
}
// UpdateIAM2Organization updates IAM2Organization
func (cli *ZSClient) UpdateIAM2Organization(uuid string, params param.UpdateIAM2OrganizationParam) (*view.IAM2OrganizationInventoryView, error) {
	resp := view.IAM2OrganizationInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/organizations", uuid, "", map[string]interface{}{
		"updateIAM2Organization": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateIAM2Organization creates IAM2Organization
func (cli *ZSClient) CreateIAM2Organization(params param.CreateIAM2OrganizationParam) (*view.IAM2OrganizationInventoryView, error) {
	resp := view.IAM2OrganizationInventoryView{}
	if err := cli.Post("v1/iam2/organizations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeIAM2OrganizationState changes IAM2OrganizationState
func (cli *ZSClient) ChangeIAM2OrganizationState(uuid string, params param.ChangeIAM2OrganizationStateParam) (*view.IAM2OrganizationInventoryView, error) {
	resp := view.IAM2OrganizationInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/organizations", uuid, "", map[string]interface{}{
		"changeIAM2OrganizationState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2OrganizationVirtualIDNumber gets IAM2OrganizationVirtualIDNumber by uuid
func (cli *ZSClient) GetIAM2OrganizationVirtualIDNumber(uuid string) (*view.GetIAM2OrganizationVirtualIDNumberView, error) {
	var resp view.GetIAM2OrganizationVirtualIDNumberView
	if err := cli.GetWithRespKey("v1/iam2/organizations", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetOrganizationSupervisor operates on OrganizationSupervisor
func (cli *ZSClient) SetOrganizationSupervisor(uuid string, params param.SetOrganizationSupervisorParam) (*view.SetOrganizationSupervisorEventView, error) {
	resp := view.SetOrganizationSupervisorEventView{}
	if err := cli.PutWithRespKey("v1/iam2/organizations", uuid, "", map[string]interface{}{
		"setOrganizationSupervisor": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2SystemAttributes gets IAM2SystemAttributes by uuid
func (cli *ZSClient) GetIAM2SystemAttributes() (*view.IAM2AttributeInventoryView, error) {
	var resp view.IAM2AttributeInventoryView
	if err := cli.GetWithRespKey("v1/iam2/attributes/system", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrganizationQuotaUsage gets OrganizationQuotaUsage by uuid
func (cli *ZSClient) GetOrganizationQuotaUsage(uuid string) (*view.GetOrganizationQuotaUsageView, error) {
	var resp view.GetOrganizationQuotaUsageView
	if err := cli.GetWithRespKey("v1/iam2/organizations/quota", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIAM2OrganizationAvailability operates on IAM2OrganizationAvailability
func (cli *ZSClient) CheckIAM2OrganizationAvailability() (*view.CheckIAM2OrganizationAvailabilityView, error) {
	var resp view.CheckIAM2OrganizationAvailabilityView
	if err := cli.GetWithRespKey("v1/iam2/organizations/availabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetOrganizationOperation operates on OrganizationOperation
func (cli *ZSClient) SetOrganizationOperation(uuid string, params param.SetOrganizationOperationParam) (*view.SetOrganizationOperationEventView, error) {
	resp := view.SetOrganizationOperationEventView{}
	if err := cli.PutWithRespKey("v1/iam2/organizations", uuid, "", map[string]interface{}{
		"setOrganizationOperation": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeIAM2OrganizationParent changes IAM2OrganizationParent
func (cli *ZSClient) ChangeIAM2OrganizationParent(parentUuid string, params param.ChangeIAM2OrganizationParentParam) (*view.ChangeIAM2OrganizationParentEventView, error) {
	resp := view.ChangeIAM2OrganizationParentEventView{}
	if err := cli.PutWithRespKey("v1/iam2/organizations", parentUuid, "", map[string]interface{}{
		"changeIAM2OrganizationParent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
