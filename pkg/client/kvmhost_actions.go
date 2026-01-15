// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddKVMHost adds KVMHost
func (cli *ZSClient) AddKVMHost(params param.AddKVMHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.Post("v1/hosts/kvm", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddKVMHostAsync Async
func (cli *ZSClient) AddKVMHostAsync(params param.AddKVMHostParam) (string, error) {

	resource := "v1/hosts/kvm"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// UpdateKVMHost updates KVMHost
func (cli *ZSClient) UpdateKVMHost(uuid string, params param.UpdateKVMHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.Put("v1/hosts/kvm", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
