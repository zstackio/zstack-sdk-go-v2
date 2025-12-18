// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateCertificate 更新Certificate
func (cli *ZSClient) UpdateCertificate(uuid string, params param.UpdateCertificateParam) (*view.UpdateCertificateEventView, error) {
	resp := view.UpdateCertificateEventView{}
	if err := cli.Put("v1/certificates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

