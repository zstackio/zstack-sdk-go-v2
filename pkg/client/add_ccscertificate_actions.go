// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddCCSCertificate adds CCSCertificate
func (cli *ZSClient) AddCCSCertificate(params param.AddCCSCertificateParam) (*view.AddCCSCertificateEventView, error) {
	resp := view.AddCCSCertificateEventView{}
	if err := cli.Post("v1/crypto/ccs-certificate/add", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
