// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachCCSCertificateToUser 操作CCSCertificateToUser
func (cli *ZSClient) AttachCCSCertificateToUser(params param.AttachCCSCertificateToUserParam) (*view.AttachCCSCertificateToUserEventView, error) {
	resp := view.AttachCCSCertificateToUserEventView{}
	if err := cli.Post("v1/crypto/ccs-certificate/attach-user/{userUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

