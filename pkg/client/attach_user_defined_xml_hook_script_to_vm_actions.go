// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachUserDefinedXmlHookScriptToVm operates on UserDefinedXmlHookScriptToVm
func (cli *ZSClient) AttachUserDefinedXmlHookScriptToVm(params param.AttachUserDefinedXmlHookScriptToVmParam) (*view.AttachUserDefinedXmlHookScriptToVmEventView, error) {
	resp := view.AttachUserDefinedXmlHookScriptToVmEventView{}
	if err := cli.Post("v1/xmlhook/{xmlHookUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
