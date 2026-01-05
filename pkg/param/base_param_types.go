// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// This file contains nested types used in API request params

// AttributeParam Attribute param struct
type AttributeParam struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// IAM2FlowStructParam IAM2FlowStruct param struct
type IAM2FlowStructParam struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ApproverUuid string `json:"approverUuid,omitempty"`
	ApproverTitle string `json:"approverTitle,omitempty"`
}

// ArchitectureImageMappingParam ArchitectureImageMapping param struct
type ArchitectureImageMappingParam struct {
	CpuArchitecture string `json:"cpuArchitecture,omitempty"`
	VmImageUuid string `json:"vmImageUuid,omitempty"`
	DockerImage string `json:"dockerImage,omitempty"`
}

// UpdateDatasetStructParam UpdateDatasetStruct param struct
type UpdateDatasetStructParam struct {
	Uuid string `json:"uuid,omitempty"`
	UsageScenarios []string `json:"usageScenarios,omitempty"`
	DataType string `json:"dataType,omitempty"`
}

// TicketRequestParam TicketRequest param struct
type TicketRequestParam struct {
	RequestName string `json:"requestName,omitempty"`
	ApiName string `json:"apiName,omitempty"`
	ExecuteTimes int `json:"executeTimes,omitempty"`
	ApiBody interface{} `json:"apiBody,omitempty"`
}

// LabelParam Label param struct
type LabelParam struct {
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Op string `json:"op,omitempty"`
	Compatible bool `json:"compatible,omitempty"`
}

// PolicyStatementParam PolicyStatement param struct
type PolicyStatementParam struct {
	Name string `json:"name,omitempty"`
	Effect string `json:"effect,omitempty"`
	Principals []string `json:"principals,omitempty"`
	Actions []string `json:"actions,omitempty"`
	Resources []string `json:"resources,omitempty"`
}

// SecurityGroupRuleAOParam SecurityGroupRuleAO param struct
type SecurityGroupRuleAOParam struct {
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
	RemoteSecurityGroupUuid string `json:"remoteSecurityGroupUuid,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	SrcIpRange string `json:"srcIpRange,omitempty"`
	DstIpRange string `json:"dstIpRange,omitempty"`
	DstPortRange string `json:"dstPortRange,omitempty"`
	Action string `json:"action,omitempty"`
	StartPort int `json:"startPort,omitempty"`
	EndPort int `json:"endPort,omitempty"`
	AllowedCidr string `json:"allowedCidr,omitempty"`
}

// SecurityGroupRulePriorityAOParam SecurityGroupRulePriorityAO param struct
type SecurityGroupRulePriorityAOParam struct {
	RuleUuid string `json:"ruleUuid,omitempty"`
	Priority int `json:"priority,omitempty"`
}

// VmNicParamParam VmNicParam param struct
type VmNicParamParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Mac string `json:"mac,omitempty"`
	Ip string `json:"ip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Ip6 string `json:"ip6,omitempty"`
	Prefix6 string `json:"prefix6,omitempty"`
	Gateway6 string `json:"gateway6,omitempty"`
	MetaData string `json:"metaData,omitempty"`
	DriverType string `json:"driverType,omitempty"`
	VmNicType string `json:"VmNicType,omitempty"`
	State string `json:"state,omitempty"`
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty"`
	InboundBandwidth int64 `json:"inboundBandwidth,omitempty"`
	MultiQueueNum int `json:"multiQueueNum,omitempty"`
	IsDefaultNic bool `json:"isDefaultNic,omitempty"`
	SgUuids []string `json:"sgUuids,omitempty"`
}

// ThresholdParam Threshold param struct
type ThresholdParam struct {
	ThresholdName string `json:"thresholdName,omitempty"`
	ThresholdValue string `json:"thresholdValue,omitempty"`
	Operator string `json:"operator,omitempty"`
}

// ResourceConfigAOParam ResourceConfigAO param struct
type ResourceConfigAOParam struct {
	Category string `json:"category,omitempty"`
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// DiskAOParam DiskAO param struct
type DiskAOParam struct {
	Boot bool `json:"boot,omitempty"`
	Platform string `json:"platform,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	Size int64 `json:"size,omitempty"`
	TemplateUuid string `json:"templateUuid,omitempty"`
	DiskOfferingUuid string `json:"diskOfferingUuid,omitempty"`
	SourceType string `json:"sourceType,omitempty"`
	SourceUuid string `json:"sourceUuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
	Name string `json:"name,omitempty"`
}

// ActionParamParam ActionParam param struct
type ActionParamParam struct {
	ActionUuid string `json:"actionUuid,omitempty"`
	ActionType string `json:"actionType,omitempty"`
}

// IAM2ProjectRoleRefStructParam IAM2ProjectRoleRefStruct param struct
type IAM2ProjectRoleRefStructParam struct {
	ProjectUuid string `json:"projectUuid,omitempty"`
	GroupUuids []string `json:"groupUuids,omitempty"`
	RoleUuids []string `json:"roleUuids,omitempty"`
}

// ExtendedAttributeParam ExtendedAttribute param struct
type ExtendedAttributeParam struct {
	Type string `json:"type,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// PriceParam Price param struct
type PriceParam struct {
	ResourceName string `json:"resourceName,omitempty"`
	ResourceUnit string `json:"resourceUnit,omitempty"`
	TimeUnit string `json:"timeUnit,omitempty"`
	Price float64 `json:"price,omitempty"`
	DateInLong int64 `json:"dateInLong,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
}

