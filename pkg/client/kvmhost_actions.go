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
	if err := cli.PutWithRespKey("v1/hosts/kvm", uuid, "", map[string]interface{}{
		"updateKVMHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeHostPassword changes HostPassword
func (cli *ZSClient) ChangeHostPassword(hostUuid string, params param.ChangeHostPasswordParam) (*view.ChangeHostPasswordEventView, error) {
	resp := view.ChangeHostPasswordEventView{}
	if err := cli.PutWithRespKey("v1/hosts/kvm", hostUuid, "", map[string]interface{}{
		"changeHostPassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IdentifyHost operates on IdentifyHost
func (cli *ZSClient) IdentifyHost(uuid string, params param.IdentifyHostParam) (*view.IdentifyHostEventView, error) {
	resp := view.IdentifyHostEventView{}
	if err := cli.PutWithRespKey("v1/hosts/kvm", uuid, "", map[string]interface{}{
		"identifyHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckKVMHostConfigFile operates on KVMHostConfigFile
func (cli *ZSClient) CheckKVMHostConfigFile() (*view.CheckHostConfigFileView, error) {
	resp := view.CheckHostConfigFileView{}
	if err := cli.Post("v1/hosts/kvm/from-file/check", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddKVMHostFromConfigFile adds KVMHostFromConfigFile
func (cli *ZSClient) AddKVMHostFromConfigFile() (*view.AddHostFromConfigFileEventView, error) {
	resp := view.AddHostFromConfigFileEventView{}
	if err := cli.Post("v1/hosts/kvm/from-file", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddKVMHostFromConfigFileAsync Async
func (cli *ZSClient) AddKVMHostFromConfigFileAsync(params param.AddKVMHostFromConfigFileParam) (string, error) {

	resource := "v1/hosts/kvm/from-file"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// KvmRunShell operates on KvmRunShell
func (cli *ZSClient) KvmRunShell(params param.KvmRunShellParam) (*view.KvmRunShellEventView, error) {
	resp := view.KvmRunShellEventView{}
	if err := cli.PutWithRespKey("v1/hosts/kvm/actions", "", "", map[string]interface{}{
		"kvmRunShell": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
