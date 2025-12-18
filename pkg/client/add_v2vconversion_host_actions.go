// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddV2VConversionHost adds V2VConversionHost
func (cli *ZSClient) AddV2VConversionHost(params param.AddV2VConversionHostParam) (*view.AddV2VConversionHostEventView, error) {
	resp := view.AddV2VConversionHostEventView{}
	if err := cli.Post("v1/v2v-conversion-hosts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
