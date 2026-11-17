// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DeleteHostNetworkServiceType deletes HostNetworkServiceType
func (cli *ZSClient) DeleteHostNetworkServiceType(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts/service-types", uuid, string(deleteMode))
}

// CreateHostNetworkServiceType creates HostNetworkServiceType
func (cli *ZSClient) CreateHostNetworkServiceType(params param.CreateHostNetworkServiceTypeParam) (*view.HostNetworkLabelInventoryView, error) {
	resp := view.HostNetworkLabelInventoryView{}
	if err := cli.Post("v1/hosts/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHostNetworkServiceType updates HostNetworkServiceType
func (cli *ZSClient) UpdateHostNetworkServiceType(uuid string, params param.UpdateHostNetworkServiceTypeParam) (*view.HostNetworkLabelInventoryView, error) {
	resp := view.HostNetworkLabelInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts/service-types", uuid, "", map[string]interface{}{
		"updateHostNetworkServiceType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
