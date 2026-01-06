// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateNasMountTarget updates NasMountTarget
func (cli *ZSClient) UpdateNasMountTarget(uuid string, params param.UpdateNasMountTargetParam) (*view.NasMountTargetInventoryView, error) {
	var resp view.UpdateNasMountTargetEventView
	if err := cli.Put("v1/primary-storage/nas/mount/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryNasMountTarget queries NasMountTarget list
func (cli *ZSClient) QueryNasMountTarget(params *param.QueryParam) ([]view.NasMountTargetInventoryView, error) {
	var resp []view.NasMountTargetInventoryView
	return resp, cli.List("v1/primary-storage/nas/mount", params, &resp)
}
// DeleteNasMountTarget deletes NasMountTarget
func (cli *ZSClient) DeleteNasMountTarget(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/nas/mount/{uuid}", uuid, string(deleteMode))
}
