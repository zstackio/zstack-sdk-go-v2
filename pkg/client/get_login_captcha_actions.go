// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLoginCaptcha gets LoginCaptcha by uuid
func (cli *ZSClient) GetLoginCaptcha(uuid string) (*view.GetLoginCaptchaView, error) {
	var resp view.GetLoginCaptchaView
	if err := cli.Get("v1/login/control/captcha", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
