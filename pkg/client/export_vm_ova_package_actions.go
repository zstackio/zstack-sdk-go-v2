// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ExportVmOvaPackage operates on ExportVmOvaPackage
func (cli *ZSClient) ExportVmOvaPackage(params param.ExportVmOvaPackageParam) (*view.ExportVmOvaPackageEventView, error) {
	resp := view.ExportVmOvaPackageEventView{}
	if err := cli.Post("v1/ovf/ova-packages", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
