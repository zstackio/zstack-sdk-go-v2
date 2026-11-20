// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetVmXmlHookScript gets VmXmlHookScript by uuid
func (cli *ZSClient) GetVmXmlHookScript(uuid string) (*view.GetVmXmlHookScriptView, error) {
	var resp view.GetVmXmlHookScriptView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmUserDefinedXmlHookScript operates on VmUserDefinedXmlHookScript
func (cli *ZSClient) SetVmUserDefinedXmlHookScript(vmInstanceUuid string, params param.SetVmUserDefinedXmlHookScriptParam) (*view.SetVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.SetVmUserDefinedXmlHookScriptEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"setVmUserDefinedXmlHookScript": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmUserDefinedXmlHookScript updates VmUserDefinedXmlHookScript
func (cli *ZSClient) UpdateVmUserDefinedXmlHookScript(params param.UpdateVmUserDefinedXmlHookScriptParam) (*view.XmlHookInventoryView, error) {
	resp := view.XmlHookInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances/xml-hook-script", "", "", map[string]interface{}{
		"updateVmUserDefinedXmlHookScript": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExpungeVmUserDefinedXmlHookScript operates on VmUserDefinedXmlHookScript
func (cli *ZSClient) ExpungeVmUserDefinedXmlHookScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/xml-hook-script", uuid, string(deleteMode))
}

// DetachUserDefinedXmlHookScriptFromVm sends the action with its default options.
func (cli *ZSClient) DetachUserDefinedXmlHookScriptFromVm(uuid string, deleteMode param.DeleteMode) error {
	return cli.DetachUserDefinedXmlHookScriptFromVmWithParams(uuid, deleteMode, param.DetachUserDefinedXmlHookScriptFromVmParamDetail{})
}

// DetachUserDefinedXmlHookScriptFromVmWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DetachUserDefinedXmlHookScriptFromVmWithParams(vmInstanceUuid string, deleteMode param.DeleteMode, params param.DetachUserDefinedXmlHookScriptFromVmParamDetail) error {
	query := url.Values{}
	if params.StartupStrategy != nil {
		query.Set("startupStrategy", *params.StartupStrategy)
	}
	return cli.deleteActionWithQuery("v1/xmlhook/vm-instances", vmInstanceUuid, "detach", deleteModeOrDefault(deleteMode, params.DeleteMode), query)
}

// DeleteVmUserDefinedXmlHookScript deletes VmUserDefinedXmlHookScript
func (cli *ZSClient) DeleteVmUserDefinedXmlHookScript(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}

// QueryVmUserDefinedXmlHookScript queries VmUserDefinedXmlHookScript list
func (cli *ZSClient) QueryVmUserDefinedXmlHookScript(params *param.QueryParam) ([]view.XmlHookInventoryView, error) {
	var resp []view.XmlHookInventoryView
	return resp, cli.List("v1/vm-instances/xml-hook-script", params, &resp)
}

func (cli *ZSClient) GetVmUserDefinedXmlHookScript(uuid string) (*view.XmlHookInventoryView, error) {
	var resp view.XmlHookInventoryView
	if err := cli.Get("v1/vm-instances/xml-hook-script", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmUserDefinedXmlHookScript Pagination
func (cli *ZSClient) PageVmUserDefinedXmlHookScript(params *param.QueryParam) ([]view.XmlHookInventoryView, int, error) {
	var vmUserDefinedXmlHookScripts []view.XmlHookInventoryView
	total, err := cli.Page("v1/vm-instances/xml-hook-script", params, &vmUserDefinedXmlHookScripts)
	return vmUserDefinedXmlHookScripts, total, err
}

// CreateVmUserDefinedXmlHookScript creates VmUserDefinedXmlHookScript
func (cli *ZSClient) CreateVmUserDefinedXmlHookScript(params param.CreateVmUserDefinedXmlHookScriptParam) (*view.XmlHookInventoryView, error) {
	resp := view.XmlHookInventoryView{}
	if err := cli.Post("v1/vm-instances/xml-hook-script", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachUserDefinedXmlHookScriptToVm operates on UserDefinedXmlHookScriptToVm
func (cli *ZSClient) AttachUserDefinedXmlHookScriptToVm(xmlHookUuid string, vmInstanceUuid string, params param.AttachUserDefinedXmlHookScriptToVmParam) (*view.AttachUserDefinedXmlHookScriptToVmEventView, error) {
	resp := view.AttachUserDefinedXmlHookScriptToVmEventView{}
	if err := cli.Post(fmt.Sprintf("v1/xmlhook/%s/vm-instances/%s", xmlHookUuid, vmInstanceUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
