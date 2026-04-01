// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAliyunNasMountTarget adds AliyunNasMountTarget
func (cli *ZSClient) AddAliyunNasMountTarget(params param.AddAliyunNasMountTargetParam) (*view.AliyunNasMountTargetInventoryView, error) {
	resp := view.AliyunNasMountTargetInventoryView{}
	if err := cli.PutWithRespKey("v1/nas/aliyun/mount", "", "", map[string]interface{}{
		"addAliyunNasMountTarget": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateAliyunNasMountTarget creates AliyunNasMountTarget
func (cli *ZSClient) CreateAliyunNasMountTarget(params param.CreateAliyunNasMountTargetParam) (*view.NasMountTargetInventoryView, error) {
	resp := view.NasMountTargetInventoryView{}
	if err := cli.Post("v1/nas/aliyun/mount", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
