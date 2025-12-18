// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeV2VConversionHostState 操作V2VConversionHostState
func (cli *ZSClient) ChangeV2VConversionHostState(uuid string, params param.ChangeV2VConversionHostStateParam) (*view.ChangeV2VConversionHostStateEventView, error) {
	resp := view.ChangeV2VConversionHostStateEventView{}
	if err := cli.Put("v1/v2v-conversion-hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

