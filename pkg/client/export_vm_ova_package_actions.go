// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ExportVmOvaPackage 操作ExportVmOvaPackage
func (cli *ZSClient) ExportVmOvaPackage(params param.ExportVmOvaPackageParam) (*view.ExportVmOvaPackageEventView, error) {
	resp := view.ExportVmOvaPackageEventView{}
	if err := cli.Post("v1/ovf/ova-packages", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

