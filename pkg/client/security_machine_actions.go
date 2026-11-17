// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSecurityMachine deletes SecurityMachine
func (cli *ZSClient) DeleteSecurityMachine(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-machines", uuid, string(deleteMode))
}
// UpdateSecurityMachine updates SecurityMachine
func (cli *ZSClient) UpdateSecurityMachine(uuid string, params param.UpdateSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithRespKey("v1/security-machines", uuid, "", map[string]interface{}{
		"updateSecurityMachine": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySecurityMachine queries SecurityMachine list
func (cli *ZSClient) QuerySecurityMachine(params *param.QueryParam) ([]view.SecurityMachineInventoryView, error) {
	var resp []view.SecurityMachineInventoryView
	return resp, cli.List("v1/security-machines", params, &resp)
}

func (cli *ZSClient) GetSecurityMachine(uuid string) (*view.SecurityMachineInventoryView, error) {
	var resp view.SecurityMachineInventoryView
	if err := cli.Get("v1/security-machines", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSecurityMachine Pagination
func (cli *ZSClient) PageSecurityMachine(params *param.QueryParam) ([]view.SecurityMachineInventoryView, int, error) {
	var securityMachines []view.SecurityMachineInventoryView
	total, err := cli.Page("v1/security-machines", params, &securityMachines)
	return securityMachines, total, err
}

// SecurityMachineDetectSync operates on MachineDetectSync
func (cli *ZSClient) SecurityMachineDetectSync(uuid string, params param.SecurityMachineDetectSyncParam) (*view.SecurityMachineDetectSyncEventView, error) {
	resp := view.SecurityMachineDetectSyncEventView{}
	if err := cli.Post(fmt.Sprintf("v1/security-machine/%s/detect/sync/actions", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetSecurityMachineKey operates on SecurityMachineKey
func (cli *ZSClient) SetSecurityMachineKey(uuid string, params param.SetSecurityMachineKeyParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/secret-resource-pool-token/set/%s/actions", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SecurityMachineEncrypt operates on MachineEncrypt
func (cli *ZSClient) SecurityMachineEncrypt(params param.SecurityMachineEncryptParam) (*view.SecurityMachineEncryptEventView, error) {
	resp := view.SecurityMachineEncryptEventView{}
	if err := cli.Post("v1/security-machine/encrypt/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecurityMachineState changes SecurityMachineState
func (cli *ZSClient) ChangeSecurityMachineState(uuid string, params param.ChangeSecurityMachineStateParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithRespKey("v1/security-machines", uuid, "", map[string]interface{}{
		"changeSecurityMachineState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
