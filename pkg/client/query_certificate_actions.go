// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryCertificate queries Certificate list
func (cli *ZSClient) QueryCertificate(params *param.QueryParam) ([]view.CertificateInventoryView, error) {
	var resp []view.CertificateInventoryView
	return resp, cli.List("v1/certificates", params, &resp)
}
