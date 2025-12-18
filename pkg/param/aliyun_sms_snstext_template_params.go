// Copyright (c) ZStack.io, Inc.

package param

// UpdateAliyunSmsSNSTextTemplateDetailParam UpdateAliyunSmsSNSTextTemplate详细参数
type UpdateAliyunSmsSNSTextTemplateDetailParam struct {
	rest string `json:"alarmTemplateCode,omitempty"`
	rest string `json:"sign,omitempty"`
	rest string `json:"eventTemplateCode,omitempty"`
	rest string `json:"eventTemplate,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"subject,omitempty"`
	rest string `json:"recoverySubject,omitempty"`
	rest string `json:"template,omitempty"`
	rest string `json:"recoveryTemplate,omitempty"`
	rest bool `json:"defaultTemplate,omitempty"`
}

// UpdateAliyunSmsSNSTextTemplateParam UpdateAliyunSmsSNSTextTemplate请求参数
type UpdateAliyunSmsSNSTextTemplateParam struct {
	BaseParam
	Params UpdateAliyunSmsSNSTextTemplateDetailParam `json:"params"` // 详细参数
}

