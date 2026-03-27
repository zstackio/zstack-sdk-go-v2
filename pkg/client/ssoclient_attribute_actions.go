// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateSSOClientAttribute updates SSOClientAttribute
func (cli *ZSClient) UpdateSSOClientAttribute(ctx context.Context, uuid string, params param.UpdateSSOClientAttributeParam) (*view.SSOClientAttributeInventoryView, error) {
	resp := view.SSOClientAttributeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sso/client/attributes", uuid, "", map[string]interface{}{
		"updateSSOClientAttribute": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
