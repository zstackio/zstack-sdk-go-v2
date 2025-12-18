// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateCCSCertificateUserState 更新CCSCertificateUserState
func (cli *ZSClient) UpdateCCSCertificateUserState(uuid string, params param.UpdateCCSCertificateUserStateParam) (*view.UpdateCCSCertificateUserStateEventView, error) {
	resp := view.UpdateCCSCertificateUserStateEventView{}
	if err := cli.Put("v1/crypto/ccs-certificate/update-state/{userUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

