// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachUserDefinedXmlHookScriptToVm operates on UserDefinedXmlHookScriptToVm
func (cli *ZSClient) AttachUserDefinedXmlHookScriptToVm(params param.AttachUserDefinedXmlHookScriptToVmParam) (*view.AttachUserDefinedXmlHookScriptToVmEventView, error) {
	resp := view.AttachUserDefinedXmlHookScriptToVmEventView{}
	if err := cli.Post("v1/xmlhook/{xmlHookUuid}/vm-instances/{vmInstanceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
