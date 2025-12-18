// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmUserDefinedXmlHookScript queries VmUserDefinedXmlHookScript list
func (cli *ZSClient) QueryVmUserDefinedXmlHookScript(params param.QueryParam) ([]view.XmlHookInventoryView, error) {
	var resp []view.XmlHookInventoryView
	return resp, cli.List("v1/vm-instances/xml-hook-script", &params, &resp)
}
