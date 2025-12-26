// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachCCSCertificateToUser operates on CCSCertificateToUser
func (cli *ZSClient) AttachCCSCertificateToUser(params param.AttachCCSCertificateToUserParam) (*view.AttachCCSCertificateToUserEventView, error) {
	resp := view.AttachCCSCertificateToUserEventView{}
	if err := cli.Post("v1/crypto/ccs-certificate/attach-user/{userUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
