// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateNasMountTarget updates NasMountTarget
func (cli *ZSClient) UpdateNasMountTarget(uuid string, params param.UpdateNasMountTargetParam) (*view.NasMountTargetInventoryView, error) {
	resp := view.NasMountTargetInventoryView{}
	if err := cli.Put("v1/primary-storage/nas/mount", uuid, map[string]interface{}{
		"updateNasMountTarget": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryNasMountTarget queries NasMountTarget list
func (cli *ZSClient) QueryNasMountTarget(params *param.QueryParam) ([]view.NasMountTargetInventoryView, error) {
	var resp []view.NasMountTargetInventoryView
	return resp, cli.List("v1/primary-storage/nas/mount", params, &resp)
}

func (cli *ZSClient) GetNasMountTarget(uuid string) (*view.NasMountTargetInventoryView, error) {
	var resp view.NasMountTargetInventoryView
	if err := cli.Get("v1/primary-storage/nas/mount", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNasMountTarget Pagination
func (cli *ZSClient) PageNasMountTarget(params *param.QueryParam) ([]view.NasMountTargetInventoryView, int, error) {
	var nasMountTargets []view.NasMountTargetInventoryView
	total, err := cli.Page("v1/primary-storage/nas/mount", params, &nasMountTargets)
	return nasMountTargets, total, err
}
// DeleteNasMountTarget deletes NasMountTarget
func (cli *ZSClient) DeleteNasMountTarget(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/nas/mount", uuid, string(deleteMode))
}
