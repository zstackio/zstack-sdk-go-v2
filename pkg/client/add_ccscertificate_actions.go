// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddCCSCertificate 操作AddCCSCertificate
func (cli *ZSClient) AddCCSCertificate(uuid string, params param.AddCCSCertificateParam) (*view.AddCCSCertificateEventView, error) {
	resp := view.AddCCSCertificateEventView{}
	if err := cli.Put("v1/crypto/ccs-certificate/add", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

