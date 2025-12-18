// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RunIAM2Script operates on RunIAM2Script
func (cli *ZSClient) RunIAM2Script(params param.RunIAM2ScriptParam) (*view.RunIAM2ScriptEventView, error) {
	resp := view.RunIAM2ScriptEventView{}
	if err := cli.Post("v1/iam2/iam2-script/run", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
