// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAliyunNasMountTarget adds AliyunNasMountTarget
func (cli *ZSClient) AddAliyunNasMountTarget(params param.AddAliyunNasMountTargetParam) (*view.AliyunNasMountTargetInventoryView, error) {
	var resp view.AddAliyunNasMountTargetEventView
	if err := cli.Post("v1/nas/aliyun/mount", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateAliyunNasMountTarget creates AliyunNasMountTarget
func (cli *ZSClient) CreateAliyunNasMountTarget(params param.CreateAliyunNasMountTargetParam) (*view.NasMountTargetInventoryView, error) {
	var resp view.CreateNasMountTargetEventView
	if err := cli.Post("v1/nas/aliyun/mount", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
