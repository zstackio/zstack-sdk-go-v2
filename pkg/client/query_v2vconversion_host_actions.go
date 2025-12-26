// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryV2VConversionHost queries V2VConversionHost list
func (cli *ZSClient) QueryV2VConversionHost(params *param.QueryParam) ([]view.V2VConversionHostInventoryView, error) {
	var resp []view.V2VConversionHostInventoryView
	return resp, cli.List("v1/v2v-conversion-hosts", params, &resp)
}
