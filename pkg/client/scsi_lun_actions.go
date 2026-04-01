// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateScsiLun updates ScsiLun
func (cli *ZSClient) UpdateScsiLun(uuid string, params param.UpdateScsiLunParam) (*view.ScsiLunInventoryView, error) {
	resp := view.ScsiLunInventoryView{}
	if err := cli.PutWithRespKey("v1/storage-devices/scsi-lun", uuid, "", map[string]interface{}{
		"updateScsiLun": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryScsiLun queries ScsiLun list
func (cli *ZSClient) QueryScsiLun(params *param.QueryParam) ([]view.ScsiLunInventoryView, error) {
	var resp []view.ScsiLunInventoryView
	return resp, cli.List("v1/storage-devices/scsi-lun/luns", params, &resp)
}

func (cli *ZSClient) GetScsiLun(uuid string) (*view.ScsiLunInventoryView, error) {
	var resp view.ScsiLunInventoryView
	if err := cli.Get("v1/storage-devices/scsi-lun/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageScsiLun Pagination
func (cli *ZSClient) PageScsiLun(params *param.QueryParam) ([]view.ScsiLunInventoryView, int, error) {
	var scsiLuns []view.ScsiLunInventoryView
	total, err := cli.Page("v1/storage-devices/scsi-lun/luns", params, &scsiLuns)
	return scsiLuns, total, err
}
