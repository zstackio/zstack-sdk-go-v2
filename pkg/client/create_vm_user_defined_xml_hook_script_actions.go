// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVmUserDefinedXmlHookScript creates VmUserDefinedXmlHookScript
func (cli *ZSClient) CreateVmUserDefinedXmlHookScript(params param.CreateVmUserDefinedXmlHookScriptParam) (*view.CreateVmUserDefinedXmlHookScriptEventView, error) {
	resp := view.CreateVmUserDefinedXmlHookScriptEventView{}
	if err := cli.Post("v1/vm-instances/xml-hook-script", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
