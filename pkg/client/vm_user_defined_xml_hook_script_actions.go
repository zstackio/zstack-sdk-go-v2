// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVmUserDefinedXmlHookScript 创建VmUserDefinedXmlHookScript
func (cli *ZSClient) CreateVmUserDefinedXmlHookScript(params param.CreateVmUserDefinedXmlHookScriptParam) (*view.CreateVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.CreateVmUserDefinedXmlHookScriptEventView{}
	if err := cli.Post("v1/vm-instances/xml-hook-script", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

