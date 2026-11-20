// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/url"
	"strconv"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/errors"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVipQos sends the action with its default options.
func (cli *ZSClient) DeleteVipQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteVipQosWithParams(uuid, deleteMode, param.DeleteVipQosParamDetail{})
}

// DeleteVipQosWithParams sends the required route and query parameters for DeleteVipQos.
func (cli *ZSClient) DeleteVipQosWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteVipQosParamDetail) error {
	query := url.Values{}
	if params.Port != nil {
		query.Set("port", strconv.Itoa(*params.Port))
	}
	return cli.deleteActionWithQuery("v1/vips", uuid, "vip-qos", deleteMode, query)
}
// SetVipQos operates on VipQos
func (cli *ZSClient) SetVipQos(uuid string, params param.SetVipQosParam) (*view.VipQosInventoryView, error) {
	resp := view.VipQosInventoryView{}
	if err := cli.PutWithRespKey("v1/vips", uuid, "", map[string]interface{}{
		"setVipQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// GetVipQos gets the default QoS rule for a VIP.
func (cli *ZSClient) GetVipQos(uuid string) (*view.VipQosInventoryView, error) {
	return cli.GetVipQosByPort(uuid, 0)
}

// GetVipQosByPort gets the QoS rule for one VIP and port.
func (cli *ZSClient) GetVipQosByPort(vipUuid string, port int) (*view.VipQosInventoryView, error) {
	vip, err := actionPathSegment("vipUuid", vipUuid)
	if err != nil {
		return nil, err
	}
	var resp view.GetVipQosView
	if err := cli.GetWithSpec("v1/vip", vip, "vip-qos", "", nil, &resp); err != nil {
		return nil, err
	}
	for _, rule := range resp.Inventories {
		if rule.Port == port {
			return &rule, nil
		}
	}
	return nil, errors.ErrNotFound
}
