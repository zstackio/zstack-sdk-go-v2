// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetSpiceCertificates gets SpiceCertificates by uuid
func (cli *ZSClient) GetSpiceCertificates(uuid string) (*view.GetSpiceCertificatesView, error) {
	var resp view.GetSpiceCertificatesView
	if err := cli.Get("v1/spice/certificates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
