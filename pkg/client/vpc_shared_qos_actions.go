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

// CreateVpcSharedQos creates VpcSharedQos
func (cli *ZSClient) CreateVpcSharedQos(params param.CreateVpcSharedQosParam) (*view.VpcSharedQosInventoryView, error) {
	resp := view.VpcSharedQosInventoryView{}
	if err := cli.Post("v1/vips/sharedqos", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVpcSharedQos queries VpcSharedQos list
func (cli *ZSClient) QueryVpcSharedQos(params *param.QueryParam) ([]view.VpcSharedQosInventoryView, error) {
	var resp []view.VpcSharedQosInventoryView
	return resp, cli.List("v1/vips/sharedqos", params, &resp)
}

func (cli *ZSClient) GetVpcSharedQos(uuid string) (*view.VpcSharedQosInventoryView, error) {
	var resp view.VpcSharedQosInventoryView
	if err := cli.Get("v1/vips/sharedqos", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcSharedQos Pagination
func (cli *ZSClient) PageVpcSharedQos(params *param.QueryParam) ([]view.VpcSharedQosInventoryView, int, error) {
	var vpcSharedQos []view.VpcSharedQosInventoryView
	total, err := cli.Page("v1/vips/sharedqos", params, &vpcSharedQos)
	return vpcSharedQos, total, err
}
// UpdateVpcSharedQos updates VpcSharedQos
func (cli *ZSClient) UpdateVpcSharedQos(uuid string, params param.UpdateVpcSharedQosParam) (*view.VpcSharedQosInventoryView, error) {
	resp := view.VpcSharedQosInventoryView{}
	if err := cli.PutWithRespKey("v1/vips/sharedqos", uuid, "", map[string]interface{}{
		"updateVpcSharedQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVpcSharedQos deletes VpcSharedQos
func (cli *ZSClient) DeleteVpcSharedQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips/sharedqos", uuid, string(deleteMode))
}

// AttachVipToVpcSharedQos operates on VipToVpcSharedQos
func (cli *ZSClient) AttachVipToVpcSharedQos(sharedQosUuid string, params param.AttachVipToVpcSharedQosParam) (*view.AttachVipToVpcSharedQosEventView, error) {
	resp := view.AttachVipToVpcSharedQosEventView{}
	if err := cli.Post(fmt.Sprintf("v1/vips/sharedqos/%s/vips", sharedQosUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVpcSharedQosBandwidth changes VpcSharedQosBandwidth
func (cli *ZSClient) ChangeVpcSharedQosBandwidth(sharedQosUuid string, params param.ChangeVpcSharedQosBandwidthParam) (*view.VpcSharedQosInventoryView, error) {
	resp := view.VpcSharedQosInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vips/sharedqos/%s/bandwidth/actions", sharedQosUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachVipFromVpcSharedQos requires the VIPs to detach.
//
// Deprecated: use DetachVipFromVpcSharedQosWithParams to specify the VIPs.
func (cli *ZSClient) DetachVipFromVpcSharedQos(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DetachVipFromVpcSharedQos requires VIP UUIDs; use DetachVipFromVpcSharedQosWithParams")
}

// DetachVipFromVpcSharedQosWithParams detaches VIPs identified by VipUuids or VipLists.
func (cli *ZSClient) DetachVipFromVpcSharedQosWithParams(sharedQosUuid string, deleteMode param.DeleteMode, params param.DetachVipFromVpcSharedQosParamDetail) error {
	query := url.Values{}
	var required []string
	if len(params.VipUuids) > 0 {
		query["vipUuids"] = params.VipUuids
		required = append(required, "vipUuids")
	}
	if len(params.VipLists) > 0 {
		query["vipLists"] = params.VipLists
		required = append(required, "vipLists")
	}
	if len(required) == 0 {
		return fmt.Errorf("vipUuids or vipLists is required")
	}
	return cli.deleteActionWithQuery("v1/vips/sharedqos", sharedQosUuid, "vips", deleteMode, query, required...)
}
