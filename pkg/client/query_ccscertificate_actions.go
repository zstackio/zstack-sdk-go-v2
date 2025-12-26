// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryCCSCertificate queries CCSCertificate list
func (cli *ZSClient) QueryCCSCertificate(params *param.QueryParam) ([]view.CCSCertificateInventoryView, error) {
	var resp []view.CCSCertificateInventoryView
	return resp, cli.List("v1/crypto/ccs-certificate/certificates/", params, &resp)
}
