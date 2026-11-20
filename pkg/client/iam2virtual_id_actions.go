// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2VirtualID queries IAM2VirtualID list
func (cli *ZSClient) QueryIAM2VirtualID(params *param.QueryParam) ([]view.IAM2VirtualIDInventoryView, error) {
	var resp []view.IAM2VirtualIDInventoryView
	return resp, cli.List("v1/iam2/virtual-ids", params, &resp)
}

func (cli *ZSClient) GetIAM2VirtualID(uuid string) (*view.IAM2VirtualIDInventoryView, error) {
	var resp view.IAM2VirtualIDInventoryView
	if err := cli.Get("v1/iam2/virtual-ids", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2VirtualID Pagination
func (cli *ZSClient) PageIAM2VirtualID(params *param.QueryParam) ([]view.IAM2VirtualIDInventoryView, int, error) {
	var iAM2VirtualIDs []view.IAM2VirtualIDInventoryView
	total, err := cli.Page("v1/iam2/virtual-ids", params, &iAM2VirtualIDs)
	return iAM2VirtualIDs, total, err
}
// CreateIAM2VirtualID creates IAM2VirtualID
func (cli *ZSClient) CreateIAM2VirtualID(params param.CreateIAM2VirtualIDParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.Post("v1/iam2/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIAM2VirtualID deletes IAM2VirtualID
func (cli *ZSClient) DeleteIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/virtual-ids", uuid, string(deleteMode))
}
// LoginIAM2VirtualID operates on IAM2VirtualID
func (cli *ZSClient) LoginIAM2VirtualID(params param.LoginIAM2VirtualIDParam) (*view.SessionInventoryView, error) {
	resp := view.SessionInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/virtual-ids/login", "", "", map[string]interface{}{
		"loginIAM2VirtualID": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIAM2VirtualID updates IAM2VirtualID
func (cli *ZSClient) UpdateIAM2VirtualID(uuid string, params param.UpdateIAM2VirtualIDParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/virtual-ids", uuid, "", map[string]interface{}{
		"updateIAM2VirtualID": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchCreateIAM2VirtualIDFromConfigFile operates on CreateIAM2VirtualIDFromConfigFile
func (cli *ZSClient) BatchCreateIAM2VirtualIDFromConfigFile(params param.BatchCreateIAM2VirtualIDFromConfigFileParam) (*view.BatchCreateIAM2VirtualIDFromConfigFileEventView, error) {
	resp := view.BatchCreateIAM2VirtualIDFromConfigFileEventView{}
	if err := cli.Post("v1/iam2/virtual-ids/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIAM2VirtualIDsToOrganization adds IAM2VirtualIDsToOrganization
func (cli *ZSClient) AddIAM2VirtualIDsToOrganization(organizationUuid string, params param.AddIAM2VirtualIDsToOrganizationParam) (*view.AddIAM2VirtualIDsToOrganizationEventView, error) {
	resp := view.AddIAM2VirtualIDsToOrganizationEventView{}
	if err := cli.Post(fmt.Sprintf("v1/iam2/organizations/%s/virtual-ids", organizationUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIAM2VirtualIDsToProject adds IAM2VirtualIDsToProject
func (cli *ZSClient) AddIAM2VirtualIDsToProject(projectUuid string, params param.AddIAM2VirtualIDsToProjectParam) (*view.AddIAM2VirtualIDsToProjectEventView, error) {
	resp := view.AddIAM2VirtualIDsToProjectEventView{}
	if err := cli.Post(fmt.Sprintf("v1/iam2/projects/%s/virtual-ids", projectUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeIAM2VirtualIDType changes IAM2VirtualIDType
func (cli *ZSClient) ChangeIAM2VirtualIDType(uuid string, params param.ChangeIAM2VirtualIDTypeParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/virtual-ids", uuid, "", map[string]interface{}{
		"changeIAM2VirtualIDType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddIAM2VirtualIDsToProjects adds IAM2VirtualIDsToProjects
func (cli *ZSClient) AddIAM2VirtualIDsToProjects(params param.AddIAM2VirtualIDsToProjectsParam) (*view.AddIAM2VirtualIDsToProjectsEventView, error) {
	resp := view.AddIAM2VirtualIDsToProjectsEventView{}
	if err := cli.Post("v1/iam2/projects/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRolesToIAM2VirtualID adds RolesToIAM2VirtualID
func (cli *ZSClient) AddRolesToIAM2VirtualID(virtualIDUuid string, params param.AddRolesToIAM2VirtualIDParam) (*view.AddRolesToIAM2VirtualIDEventView, error) {
	resp := view.AddRolesToIAM2VirtualIDEventView{}
	if err := cli.Post(fmt.Sprintf("v1/iam2/projects/virtual-ids/%s/roles", virtualIDUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanInvalidLdapIAM2Binding operates on InvalidLdapIAM2Binding
func (cli *ZSClient) CleanInvalidLdapIAM2Binding() (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/ldap/bindings/actions", "", "", map[string]interface{}{
		"cleanInvalidLdapIAM2Binding": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2VirtualIDInGroup gets IAM2VirtualIDInGroup by uuid
func (cli *ZSClient) GetIAM2VirtualIDInGroup() (*view.IAM2VirtualIDInventoryView, error) {
	var resp view.IAM2VirtualIDInventoryView
	if err := cli.GetWithRespKey("v1/iam2/IAM2VirtualIDGroup/IAM2VirtualID", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIAM2VirtualIDAPIPermission gets IAM2VirtualIDAPIPermission by uuid
func (cli *ZSClient) GetIAM2VirtualIDAPIPermission() (*view.GetIAM2VirtualIDAPIPermissionView, error) {
	var resp view.GetIAM2VirtualIDAPIPermissionView
	if err := cli.GetWithRespKey("v1/iam2/virtual-ids/api-permissions", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveIAM2VirtualIDsFromProjects requires additional arguments.
//
// Deprecated: use RemoveIAM2VirtualIDsFromProjectsWithParams to supply the required fields.
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProjects(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveIAM2VirtualIDsFromProjects requires additional arguments; use RemoveIAM2VirtualIDsFromProjectsWithParams")
}

// RemoveIAM2VirtualIDsFromProjectsWithParams sends the required route and query parameters for RemoveIAM2VirtualIDsFromProjects.
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProjectsWithParams(deleteMode param.DeleteMode, params param.RemoveIAM2VirtualIDsFromProjectsParamDetail) error {
	query := url.Values{
		"projectUuids":   params.ProjectUuids,
		"virtualIDUuids": params.VirtualIDUuids,
	}
	return cli.deleteActionWithQuery("v1/iam2/projects/virtual-ids", "", "", deleteMode, query, "projectUuids", "virtualIDUuids")
}

// RemoveIAM2VirtualIDsFromProject requires additional arguments.
//
// Deprecated: use RemoveIAM2VirtualIDsFromProjectWithParams to supply the required fields.
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProject(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveIAM2VirtualIDsFromProject requires additional arguments; use RemoveIAM2VirtualIDsFromProjectWithParams")
}

// RemoveIAM2VirtualIDsFromProjectWithParams sends the required route and query parameters for RemoveIAM2VirtualIDsFromProject.
func (cli *ZSClient) RemoveIAM2VirtualIDsFromProjectWithParams(projectUuid string, deleteMode param.DeleteMode, params param.RemoveIAM2VirtualIDsFromProjectParamDetail) error {
	query := url.Values{
		"virtualIDUuids": params.VirtualIDUuids,
	}
	return cli.deleteActionWithQuery("v1/iam2/projects", projectUuid, "virtual-ids", deleteMode, query, "virtualIDUuids")
}

// CreateIAM2VirtualIDFromLdapUid creates IAM2VirtualIDFromLdapUid
func (cli *ZSClient) CreateIAM2VirtualIDFromLdapUid(params param.CreateIAM2VirtualIDFromLdapUidParam) (*view.LdapResourceRefInventoryView, error) {
	resp := view.LdapResourceRefInventoryView{}
	if err := cli.Post("v1/iam2/virtual-id/ldap/uid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveRolesFromIAM2VirtualID requires additional arguments.
//
// Deprecated: use RemoveRolesFromIAM2VirtualIDWithParams to supply the required fields.
func (cli *ZSClient) RemoveRolesFromIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveRolesFromIAM2VirtualID requires additional arguments; use RemoveRolesFromIAM2VirtualIDWithParams")
}

// RemoveRolesFromIAM2VirtualIDWithParams sends the required route and query parameters for RemoveRolesFromIAM2VirtualID.
func (cli *ZSClient) RemoveRolesFromIAM2VirtualIDWithParams(virtualIDUuid string, deleteMode param.DeleteMode, params param.RemoveRolesFromIAM2VirtualIDParamDetail) error {
	query := url.Values{
		"roleUuids": params.RoleUuids,
	}
	if params.ProjectUuid != nil {
		query.Set("projectUuid", *params.ProjectUuid)
	}
	return cli.deleteActionWithQuery("v1/iam2/projects/virtual-ids", virtualIDUuid, "roles", deleteMode, query, "roleUuids")
}

// RemoveIAM2VirtualIDsFromOrganization requires additional arguments.
//
// Deprecated: use RemoveIAM2VirtualIDsFromOrganizationWithParams to supply the required fields.
func (cli *ZSClient) RemoveIAM2VirtualIDsFromOrganization(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveIAM2VirtualIDsFromOrganization requires additional arguments; use RemoveIAM2VirtualIDsFromOrganizationWithParams")
}

// RemoveIAM2VirtualIDsFromOrganizationWithParams sends the required route and query parameters for RemoveIAM2VirtualIDsFromOrganization.
func (cli *ZSClient) RemoveIAM2VirtualIDsFromOrganizationWithParams(organizationUuid string, deleteMode param.DeleteMode, params param.RemoveIAM2VirtualIDsFromOrganizationParamDetail) error {
	query := url.Values{
		"virtualIDUuids": params.VirtualIDUuids,
	}
	return cli.deleteActionWithQuery("v1/iam2/organizations", organizationUuid, "virtual-ids", deleteMode, query, "virtualIDUuids")
}

// ChangeIAM2VirtualIDState changes IAM2VirtualIDState
func (cli *ZSClient) ChangeIAM2VirtualIDState(uuid string, params param.ChangeIAM2VirtualIDStateParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/virtual-ids", uuid, "", map[string]interface{}{
		"changeIAM2VirtualIDState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIAM2VirtualIDConfigFile operates on IAM2VirtualIDConfigFile
func (cli *ZSClient) CheckIAM2VirtualIDConfigFile(params param.CheckIAM2VirtualIDConfigFileParam) (*view.CheckIAM2VirtualIDConfigFileView, error) {
	resp := view.CheckIAM2VirtualIDConfigFileView{}
	if err := cli.PutWithRespKey("v1/iam2/virtual-ids/from-file", "", "", map[string]interface{}{
		"checkIAM2VirtualIDConfigFile": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
