// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SetupExternalAlertIntegrationView SetupExternalAlertIntegration response
type SetupExternalAlertIntegrationView struct {
	EndpointUuid string `json:"endpointUuid"` // 创建或更新的 HTTP Endpoint UUID
	TopicUuid    string `json:"topicUuid"`    // 系统告警 Topic UUID
	Created      bool   `json:"created"`      // true 表示新建，false 表示更新已有
}

// CheckExternalAlertIntegrationView CheckExternalAlertIntegration response
type CheckExternalAlertIntegrationView struct {
	Healthy        bool `json:"healthy"`        // 所有检查项均通过
	EndpointExists bool `json:"endpointExists"` // Endpoint 是否存在
	UrlMatched     bool `json:"urlMatched"`     // Webhook 地址是否与期望一致
	Subscribed     bool `json:"subscribed"`     // 是否已订阅系统告警 Topic
	TemplateExists bool `json:"templateExists"` // 事件推送模板是否存在
}

// RemoveExternalAlertIntegrationView RemoveExternalAlertIntegration response
type RemoveExternalAlertIntegrationView struct {
	Success bool `json:"success"` // 操作是否成功
}
