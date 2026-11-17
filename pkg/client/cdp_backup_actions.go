// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// RevertVmFromCdpBackup operates on VmFromCdpBackup
func (cli *ZSClient) RevertVmFromCdpBackup(vmInstanceUuid string, params param.RevertVmFromCdpBackupParam) (*view.RevertVmFromCdpBackupEventView, error) {
	resp := view.RevertVmFromCdpBackupEventView{}
	if err := cli.PutWithRespKey("v1/cdp-backups", vmInstanceUuid, "", map[string]interface{}{
		"revertVmFromCdpBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVmFromCdpBackupAsync Async
func (cli *ZSClient) RevertVmFromCdpBackupAsync(vmInstanceUuid string, params param.RevertVmFromCdpBackupParam) (string, error) {

	resource := fmt.Sprintf("v1/cdp-backups/%s/actions", vmInstanceUuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// CreateVmFromCdpBackup creates VmFromCdpBackup
func (cli *ZSClient) CreateVmFromCdpBackup(params param.CreateVmFromCdpBackupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/cdp-backups/actions", "", "", map[string]interface{}{
		"createVmFromCdpBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmFromCdpBackupAsync Async
func (cli *ZSClient) CreateVmFromCdpBackupAsync(params param.CreateVmFromCdpBackupParam) (string, error) {

	resource := "v1/cdp-backups/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
