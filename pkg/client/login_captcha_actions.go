// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLoginCaptcha 获取LoginCaptcha详情
func (cli *ZSClient) GetLoginCaptcha(uuid string) (*view.GetLoginCaptchaView, error) {
	var resp view.GetLoginCaptchaView
	if err := cli.Get("v1/login/control/captcha", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

