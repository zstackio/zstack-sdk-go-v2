// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAliyunNasMountTarget adds AliyunNasMountTarget
func (cli *ZSClient) AddAliyunNasMountTarget(ctx context.Context, params param.AddAliyunNasMountTargetParam) (*view.AliyunNasMountTargetInventoryView, error) {
	resp := view.AliyunNasMountTargetInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/nas/aliyun/mount", "", "", map[string]interface{}{
		"addAliyunNasMountTarget": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateAliyunNasMountTarget creates AliyunNasMountTarget
func (cli *ZSClient) CreateAliyunNasMountTarget(ctx context.Context, params param.CreateAliyunNasMountTargetParam) (*view.NasMountTargetInventoryView, error) {
	resp := view.NasMountTargetInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/nas/aliyun/mount", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
