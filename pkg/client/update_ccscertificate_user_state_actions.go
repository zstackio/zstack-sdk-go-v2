// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateCCSCertificateUserState updates CCSCertificateUserState
func (cli *ZSClient) UpdateCCSCertificateUserState(uuid string, params param.UpdateCCSCertificateUserStateParam) (*view.UpdateCCSCertificateUserStateEventView, error) {
	resp := view.UpdateCCSCertificateUserStateEventView{}
	if err := cli.Put("v1/crypto/ccs-certificate/update-state/{userUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
