// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

type DeleteMode string

const (
	DeleteModePermissive DeleteMode = "Permissive"
	DeleteModeEnforcing  DeleteMode = "Enforcing"
)

type BaseParam struct {
	SystemTags []string `json:"systemTags,omitempty"` // System tags
	UserTags   []string `json:"userTags,omitempty"`   // User tags
	RequestIp  string   `json:"requestIp,omitempty"`  // Request IP
}

type HqlParam struct {
	OperationName string    `json:"operationName"` // Request name
	Query         string    `json:"query"`         // Query statement
	Variables     Variables `json:"variables"`     // Parameters for the statement
}

type Variables struct {
	Conditions      []Condition            `json:"conditions"`      // Conditions
	ExtraConditions []Condition            `json:"extraConditions"` // Extra conditions
	Input           map[string]interface{} `json:"input"`           // Input parameters
	PageVar         `json:",inline,omitempty"`
	Type            string `json:"type"` // Type
}

type Condition struct {
	Key   string `json:"key"`   // Key
	Op    string `json:"op"`    // Operator
	Value string `json:"value"` // Value
}

type PageVar struct {
	Start int `json:"start,omitempty"` // Start page
	Limit int `json:"limit,omitempty"` // Limit per page
}

// ========== Dynamically Generated Param Types ==========

// PolicyStatementParam PolicyStatement param struct
type PolicyStatementParam struct {
	Name string `json:"name,omitempty"`
	Effect string `json:"effect,omitempty"`
	Principals []string `json:"principals,omitempty"`
	Actions []string `json:"actions,omitempty"`
	Resources []string `json:"resources,omitempty"`
}

// LabelParam Label param struct
type LabelParam struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	Op string `json:"op,omitempty"`
	Compatible bool `json:"compatible,omitempty"`
}

// AddIAM2VirtualIDGroupToProjects_IAM2ProjectRoleRefStructParam IAM2ProjectRoleRefStruct param struct
type AddIAM2VirtualIDGroupToProjects_IAM2ProjectRoleRefStructParam struct {
	ProjectUuid *string `json:"projectUuid,omitempty"`
	GroupUuids []string `json:"groupUuids,omitempty"`
	RoleUuids []string `json:"roleUuids,omitempty"`
}

// CreatePriceTable_PriceParam Price param struct
type CreatePriceTable_PriceParam struct {
	ResourceName *string `json:"resourceName,omitempty"`
	ResourceUnit *string `json:"resourceUnit,omitempty"`
	TimeUnit *string `json:"timeUnit,omitempty"`
	Price float64 `json:"price,omitempty"`
	DateInLong *int64 `json:"dateInLong,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
}

// UpdateIAM2TicketFlowCollection_IAM2FlowStructParam IAM2FlowStruct param struct
type UpdateIAM2TicketFlowCollection_IAM2FlowStructParam struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ApproverUuid *string `json:"approverUuid,omitempty"`
	ApproverTitle *string `json:"approverTitle,omitempty"`
}

// AddSecurityGroupRule_SecurityGroupRuleAOParam SecurityGroupRuleAO param struct
type AddSecurityGroupRule_SecurityGroupRuleAOParam struct {
	Type string `json:"type,omitempty"`
	State *string `json:"state,omitempty"`
	Description *string `json:"description,omitempty"`
	RemoteSecurityGroupUuid *string `json:"remoteSecurityGroupUuid,omitempty"`
	IpVersion *int `json:"ipVersion,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	SrcIpRange *string `json:"srcIpRange,omitempty"`
	DstIpRange *string `json:"dstIpRange,omitempty"`
	DstPortRange *string `json:"dstPortRange,omitempty"`
	Action *string `json:"action,omitempty"`
	StartPort *int `json:"startPort,omitempty"`
	EndPort *int `json:"endPort,omitempty"`
	AllowedCidr *string `json:"allowedCidr,omitempty"`
}

// CreateIAM2TickFlowCollection_IAM2FlowStructParam IAM2FlowStruct param struct
type CreateIAM2TickFlowCollection_IAM2FlowStructParam struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ApproverUuid *string `json:"approverUuid,omitempty"`
	ApproverTitle *string `json:"approverTitle,omitempty"`
}

// ModelServiceParam ModelService param struct
type ModelServiceParam struct {
	NodeRank *int `json:"nodeRank,omitempty"`
	TensorParallelSize *int `json:"tensorParallelSize,omitempty"`
	IsInitialNode *bool `json:"isInitialNode,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description *string `json:"description,omitempty"`
	ModelUuid *string `json:"modelUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	VmImageUuid *string `json:"vmImageUuid,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	DatasetUuids []string `json:"datasetUuids,omitempty"`
	ModelServiceGroupUuids []string `json:"modelServiceGroupUuids,omitempty"`
	DockerImage *string `json:"dockerImage,omitempty"`
	CpuNum *int `json:"cpuNum,omitempty"`
	RequestCpuNum *int `json:"requestCpuNum,omitempty"`
	Name string `json:"name,omitempty"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	StartupParameters map[string]string `json:"startupParameters,omitempty"`
	Type string `json:"type,omitempty"`
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	MemorySize *int64 `json:"memorySize,omitempty"`
	RequestMemorySize *int64 `json:"requestMemorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	ServiceBootUptime *int `json:"serviceBootUptime,omitempty"`
	ServiceLivez *string `json:"serviceLivez,omitempty"`
	ServiceReadyz *string `json:"serviceReadyz,omitempty"`
	RootDiskOfferingUuid *string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize *int64 `json:"rootDiskSize,omitempty"`
	ProjectUuid *string `json:"projectUuid,omitempty"`
	Session SessionParam `json:"session,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
	Timeout int64 `json:"timeout,omitempty"`
}

// VmNicParamParam VmNicParam param struct
type VmNicParamParam struct {
	L3NetworkUuid *string `json:"l3NetworkUuid,omitempty"`
	Mac *string `json:"mac,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	Ip6 *string `json:"ip6,omitempty"`
	Prefix6 *string `json:"prefix6,omitempty"`
	Gateway6 *string `json:"gateway6,omitempty"`
	MetaData *string `json:"metaData,omitempty"`
	DriverType *string `json:"driverType,omitempty"`
	VmNicType *string `json:"VmNicType,omitempty"`
	State *string `json:"state,omitempty"`
	OutboundBandwidth *int64 `json:"outboundBandwidth,omitempty"`
	InboundBandwidth *int64 `json:"inboundBandwidth,omitempty"`
	MultiQueueNum *int `json:"multiQueueNum,omitempty"`
	IsDefaultNic *bool `json:"isDefaultNic,omitempty"`
	SgUuids []string `json:"sgUuids,omitempty"`
}

// VolumeFilterInfoParam VolumeFilterInfo param struct
type VolumeFilterInfoParam struct {
	DeviceId *string `json:"deviceId,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Skip bool `json:"skip,omitempty"`
}

// MetricDatumParam MetricDatum param struct
type MetricDatumParam struct {
	MetricName *string `json:"metricName,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Time *int64 `json:"time,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

// UpdateResourceConfigs_ResourceConfigAOParam ResourceConfigAO param struct
type UpdateResourceConfigs_ResourceConfigAOParam struct {
	Category *string `json:"category,omitempty"`
	Name string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// ArchitectureImageMappingParam ArchitectureImageMapping param struct
type ArchitectureImageMappingParam struct {
	CpuArchitecture *string `json:"cpuArchitecture,omitempty"`
	VmImageUuid *string `json:"vmImageUuid,omitempty"`
	DockerImage *string `json:"dockerImage,omitempty"`
}

// UpdateDatasets_UpdateDatasetStructParam UpdateDatasetStruct param struct
type UpdateDatasets_UpdateDatasetStructParam struct {
	Uuid string `json:"uuid,omitempty"`
	UsageScenarios []string `json:"usageScenarios,omitempty"`
	DataType *string `json:"dataType,omitempty"`
}

// ExtendedAttributeParam ExtendedAttribute param struct
type ExtendedAttributeParam struct {
	Type *string `json:"type,omitempty"`
	Purpose *string `json:"purpose,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// CreateVmInstance_DiskAOParam DiskAO param struct
type CreateVmInstance_DiskAOParam struct {
	Boot bool `json:"boot,omitempty"`
	Platform *string `json:"platform,omitempty"`
	GuestOsType *string `json:"guestOsType,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	PrimaryStorageUuid *string `json:"primaryStorageUuid,omitempty"`
	Size int64 `json:"size,omitempty"`
	TemplateUuid *string `json:"templateUuid,omitempty"`
	DiskOfferingUuid *string `json:"diskOfferingUuid,omitempty"`
	SourceType *string `json:"sourceType,omitempty"`
	SourceUuid *string `json:"sourceUuid,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
	Name string `json:"name,omitempty"`
}

// TicketRequestParam TicketRequest param struct
type TicketRequestParam struct {
	RequestName *string `json:"requestName,omitempty"`
	ApiName *string `json:"apiName,omitempty"`
	ExecuteTimes int `json:"executeTimes,omitempty"`
	ApiBody interface{} `json:"apiBody,omitempty"`
}

// ThresholdParam Threshold param struct
type ThresholdParam struct {
	ThresholdName *string `json:"thresholdName,omitempty"`
	ThresholdValue *string `json:"thresholdValue,omitempty"`
	Operator *string `json:"operator,omitempty"`
}

// UpdateSecurityGroupRulePriority_SecurityGroupRulePriorityAOParam SecurityGroupRulePriorityAO param struct
type UpdateSecurityGroupRulePriority_SecurityGroupRulePriorityAOParam struct {
	RuleUuid string `json:"ruleUuid,omitempty"`
	Priority int `json:"priority,omitempty"`
}

// MiniHostInfoParam MiniHostInfo param struct
type MiniHostInfoParam struct {
	Sn *string `json:"sn,omitempty"`
	DnsAddresses []string `json:"dnsAddresses,omitempty"`
	Ipmi MiniNetworkConfigStructParam `json:"ipmi,omitempty"`
	Mgmt MiniNetworkConfigStructParam `json:"mgmt,omitempty"`
}

// SetVmNicSecurityGroup_VmNicSecurityGroupRefAOParam VmNicSecurityGroupRefAO param struct
type SetVmNicSecurityGroup_VmNicSecurityGroupRefAOParam struct {
	SecurityGroupUuid string `json:"securityGroupUuid,omitempty"`
	Priority int `json:"priority,omitempty"`
}

// AttributeParam Attribute param struct
type AttributeParam struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// CreateAlarm_ActionParamParam ActionParam param struct
type CreateAlarm_ActionParamParam struct {
	ActionUuid *string `json:"actionUuid,omitempty"`
	ActionType *string `json:"actionType,omitempty"`
}

// SessionParam SessionInventory param struct
type SessionParam struct {
	Uuid string `json:"uuid,omitempty"`
	AccountUuid *string `json:"accountUuid,omitempty"`
	UserUuid *string `json:"userUuid,omitempty"`
	UserType *string `json:"userType,omitempty"`
	ExpiredDate time.Time `json:"expiredDate,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

// MiniNetworkConfigStructParam MiniNetworkConfigStruct param struct
type MiniNetworkConfigStructParam struct {
	Gw *string `json:"gw,omitempty"`
	Ip *string `json:"ip,omitempty"`
	Vlan *string `json:"vlan,omitempty"`
	Bond *string `json:"bond,omitempty"`
}

