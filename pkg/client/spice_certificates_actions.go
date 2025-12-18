// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetSpiceCertificates 获取SpiceCertificates详情
func (cli *ZSClient) GetSpiceCertificates(uuid string) (*view.GetSpiceCertificatesView, error) {
	var resp view.GetSpiceCertificatesView
	if err := cli.Get("v1/spice/certificates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

