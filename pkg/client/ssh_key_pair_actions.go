// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSshKeyPair creates SshKeyPair
func (cli *ZSClient) CreateSshKeyPair(params param.CreateSshKeyPairParam) (*view.SshKeyPairInventoryView, error) {
	var resp view.CreateSshKeyPairEventView
	if err := cli.Post("v1/ssh-key-pair", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// GenerateSshKeyPair operates on SshKeyPair
func (cli *ZSClient) GenerateSshKeyPair(params param.GenerateSshKeyPairParam) (*view.SshPrivateKeyPairInventoryView, error) {
	var resp view.GenerateSshKeyPairView
	if err := cli.Post("v1/ssh-key-pair/generate", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateSshKeyPair updates SshKeyPair
func (cli *ZSClient) UpdateSshKeyPair(uuid string, params param.UpdateSshKeyPairParam) (*view.SshKeyPairInventoryView, error) {
	var resp view.UpdateSshKeyPairEventView
	if err := cli.Put("v1/ssh-key-pair/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySshKeyPair queries SshKeyPair list
func (cli *ZSClient) QuerySshKeyPair(params *param.QueryParam) ([]view.SshKeyPairInventoryView, error) {
	var resp []view.SshKeyPairInventoryView
	return resp, cli.List("v1/ssh-key-pair", params, &resp)
}

func (cli *ZSClient) GetSshKeyPair(uuid string) (*view.SshKeyPairInventoryView, error) {
	var resp view.SshKeyPairInventoryView
	if err := cli.Get("v1/ssh-key-pair", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSshKeyPair deletes SshKeyPair
func (cli *ZSClient) DeleteSshKeyPair(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ssh-key-pair", uuid, string(deleteMode))
}
