// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmUserDefinedXmlHookScript queries VmUserDefinedXmlHookScript list
func (cli *ZSClient) QueryVmUserDefinedXmlHookScript(params *param.QueryParam) ([]view.XmlHookInventoryView, error) {
	var resp []view.XmlHookInventoryView
	return resp, cli.List("v1/vm-instances/xml-hook-script", params, &resp)
}
