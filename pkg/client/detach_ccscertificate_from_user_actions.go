// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DetachCCSCertificateFromUser operates on CCSCertificateFromUser
func (cli *ZSClient) DetachCCSCertificateFromUser(params param.DetachCCSCertificateFromUserParam) (*view.DetachCCSCertificateFromUserEventView, error) {
	resp := view.DetachCCSCertificateFromUserEventView{}
	if err := cli.Post("v1/crypto/ccs-certificate/detach-user/{userUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
