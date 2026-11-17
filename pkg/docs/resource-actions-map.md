# 客户端资源操作迁移清单

基于 `913fa91c`。共 1,266 个方法，目标为 `pkg/client` 下的 268 个文件。

## 计算与网络

| 目标文件 | 方法 |
|---|---|
| [auto_scaling_group_actions.go](../client/auto_scaling_group_actions.go) | `ChangeAutoScalingGroupState` |
| [auto_scaling_rule_actions.go](../client/auto_scaling_rule_actions.go) | `CreateAutoScalingGroupAddingNewInstanceRule`, `CreateAutoScalingGroupRemovalInstanceRule`, `ExecuteAutoScalingRule`, `UpdateAutoScalingGroupAddingNewInstanceRule`, `UpdateAutoScalingGroupRemovalInstanceRule` |
| [auto_scaling_template_actions.go](../client/auto_scaling_template_actions.go) | `AttachAutoScalingTemplateToGroup`, `DetachAutoScalingTemplateFromGroup` |
| [bare_metal2bonding_actions.go](../client/bare_metal2bonding_actions.go) | `AttachProvisionNicToBonding`, `CleanUpBareMetal2Bonding`, `DetachProvisionNicFromBonding` |
| [bare_metal2chassis_actions.go](../client/bare_metal2chassis_actions.go) | `ChangeBareMetal2ChassisState`, `GetBareMetal2ChassisPowerStatus`, `InspectBareMetal2ChassisByInstance`, `PowerOffBareMetal2Chassis`, `PowerOnBareMetal2Chassis`, `PowerResetBareMetal2Chassis` |
| [bare_metal2chassis_offering_actions.go](../client/bare_metal2chassis_offering_actions.go) | `ChangeBareMetal2ChassisOfferingState` |
| [bare_metal2gateway_actions.go](../client/bare_metal2gateway_actions.go) | `AttachBareMetal2GatewayToCluster`, `ChangeBareMetal2GatewayCluster`, `ChangeBareMetal2GatewayState`, `DetachBareMetal2GatewayFromCluster`, `GetBareMetal2GatewayAllocatorStrategies` |
| [bare_metal2instance_actions.go](../client/bare_metal2instance_actions.go) | `ChangeBareMetal2InstancePassword`, `CheckStaticProvisionIp`, `CreateBareMetal2InstanceFromVmBackup`, `CreateBareMetal2InstanceFromVolumeBackup`, `GetBareMetal2SupportedBootMode` |
| [bare_metal2ipmi_chassis_actions.go](../client/bare_metal2ipmi_chassis_actions.go) | `BatchAddBareMetal2IpmiChassis`, `CheckBareMetal2IpmiChassisConfigFile`, `CreateBareMetal2IpmiChassisHardwareInfo` |
| [bare_metal2provision_network_actions.go](../client/bare_metal2provision_network_actions.go) | `AttachBareMetal2ProvisionNetworkToCluster`, `ChangeBareMetal2ProvisionNetworkState`, `DetachBareMetal2ProvisionNetworkFromCluster`, `GetBareMetal2ProvisionNetworkIpAddressCapacity` |
| [baremetal_chassis_actions.go](../client/baremetal_chassis_actions.go) | `BatchCreateBaremetalChassis`, `BatchCreateBaremetalChassisAsync`, `ChangeBaremetalChassisState`, `CheckBaremetalChassisConfigFile`, `CleanUpBaremetalChassisBonding`, `GetBaremetalChassisPowerStatus`, `PowerOffBaremetalChassis`, `PowerOnBaremetalChassis`, `PowerResetBaremetalChassis` |
| [baremetal_pxe_server_actions.go](../client/baremetal_pxe_server_actions.go) | `AttachBaremetalPxeServerToCluster`, `DetachBaremetalPxeServerFromCluster` |
| [cluster_actions.go](../client/cluster_actions.go) | `ChangeClusterState`, `CreateMiniCluster`, `GetClusterHostNetworkFacts`, `UpdateClusterOS`, `UpdateClusterOSAsync`, `ValidateClusterSupportDRS` |
| [cluster_drs_actions.go](../client/cluster_drs_actions.go) | `ExecuteDRSScheduling`, `GetClusterDRSStatus` |
| [eip_actions.go](../client/eip_actions.go) | `ChangeEipState`, `GetVpcAttachedEip` |
| [fault_tolerance_vm_actions.go](../client/fault_tolerance_vm_actions.go) | `CreateFaultToleranceVmInstance`, `FailoverFaultToleranceVm`, `GetFaultToleranceVm`, `GetFaultToleranceVms`, `PageFaultToleranceVm`, `QueryFaultToleranceVm` |
| [flow_meter_actions.go](../client/flow_meter_actions.go) | `AddVRouterNetworksToFlowMeter`, `GetFlowMeterRouterId`, `GetVRouterFlowCounter`, `GetVRouterFlowMeterNetwork`, `GetVpcAttachedNetflow`, `PageVRouterFlowMeterNetwork`, `QueryVRouterFlowMeterNetwork`, `RemoveVRouterNetworksFromFlowMeter`, `SetFlowMeterRouterId` |
| [gpu_device_actions.go](../client/gpu_device_actions.go) | `GetGpuDeviceSpecCandidates` |
| [host_actions.go](../client/host_actions.go) | `AllocateHostResource`, `ChangeHostState`, `GetCpuMemoryCapacity`, `GetCpuMemoryCapacityWithParams`, `GetHostAllocatorStrategies`, `GetHostCandidatesForVmMigration`, `GetHostIommuState`, `GetHostIommuStatus`, `GetHostMultipathTopology`, `GetHostNUMATopology`, `GetHostNetworkFacts`, `GetHostPowerStatus`, `GetHostResourceAllocation`, `GetHostTask`, `GetHostWebSshUrl`, `GetHypervisorTypes`, `GetInterfaceServiceTypeStatistic`, `GetVmMigrationCandidateHosts`, `PowerOffHost`, `PowerOnHost`, `PowerResetHost`, `ShutdownHost`, `UpdateHostIommuState` |
| [host_network_bonding_actions.go](../client/host_network_bonding_actions.go) | `AttachNicToBonding`, `CreateBonding`, `DeleteBonding`, `DetachNicFromBonding`, `SetIpOnHostNetworkBonding`, `SetServiceTypeOnHostNetworkBonding`, `UpdateBonding` |
| [host_network_interface_actions.go](../client/host_network_interface_actions.go) | `GetCandidateInterfaceVlanIds`, `GetCandidateNetworkInterfaces`, `SetIpOnHostNetworkInterface`, `SetServiceTypeOnHostNetworkInterface` |
| [host_network_interface_lldp_actions.go](../client/host_network_interface_lldp_actions.go) | `ChangeHostNetworkInterfaceLldpMode` |
| [host_network_service_type_actions.go](../client/host_network_service_type_actions.go) | `CreateHostNetworkServiceType`, `DeleteHostNetworkServiceType`, `UpdateHostNetworkServiceType` |
| [host_physical_memory_actions.go](../client/host_physical_memory_actions.go) | `GetHostPhysicalMemoryFacts` |
| [host_scheduling_rule_group_actions.go](../client/host_scheduling_rule_group_actions.go) | `AddHostToHostSchedulingRuleGroup`, `DetachHostFromHostSchedulingRuleGroup` |
| [ip_address_actions.go](../client/ip_address_actions.go) | `DeleteIpAddress`, `GetIpAddress`, `PageIpAddress`, `QueryIpAddress` |
| [ip_range_actions.go](../client/ip_range_actions.go) | `AddIpRangeByNetworkCidr`, `AddIpv6Range`, `AddIpv6RangeByNetworkCidr`, `GetFreeIp` |
| [ipsec_connection_actions.go](../client/ipsec_connection_actions.go) | `AddRemoteCidrsToIPsecConnection`, `AttachL3NetworksToIPsecConnection`, `ChangeIPSecConnectionState`, `DetachL3NetworksFromIPsecConnection`, `GetIPSecConnection`, `GetVpcAttachedIpsec`, `GetVpcIPsecLog`, `PageIPSecConnection`, `QueryIPSecConnection`, `RemoveRemoteCidrsFromIPsecConnection` |
| [kvmhost_actions.go](../client/kvmhost_actions.go) | `AddKVMHostFromConfigFile`, `AddKVMHostFromConfigFileAsync`, `ChangeHostPassword`, `CheckKVMHostConfigFile`, `IdentifyHost`, `KvmRunShell` |
| [l2network_actions.go](../client/l2network_actions.go) | `AttachL2NetworkToCluster`, `CreateL2NoVlanNetwork`, `CreateL2TfNetwork`, `DetachL2NetworkFromCluster`, `GetCandidateClustersForAttachingL2Network`, `GetCandidateL2NetworksForAttachingCluster`, `GetL2NetworkTypes`, `GetVSwitchTypes`, `GetVfPciDeviceAvailableInL2Network` |
| [l2port_group_network_actions.go](../client/l2port_group_network_actions.go) | `CreateL2PortGroup` |
| [l2virtual_switch_network_actions.go](../client/l2virtual_switch_network_actions.go) | `CreateL2VirtualSwitch` |
| [l2vlan_network_actions.go](../client/l2vlan_network_actions.go) | `ChangeL2NetworkVlanId` |
| [l2vxlan_network_actions.go](../client/l2vxlan_network_actions.go) | `CreateL2HardwareVxlanNetwork`, `DeleteVxlanL2Network` |
| [l2vxlan_network_pool_actions.go](../client/l2vxlan_network_pool_actions.go) | `CreateL2HardwareVxlanNetworkPool`, `CreateVxlanPoolRemoteVtep`, `DeleteVxlanPoolRemoteVtep` |
| [l3network_actions.go](../client/l3network_actions.go) | `AddDnsToL3Network`, `AddHostRouteToL3Network`, `AddL3NetworkToGroup`, `ChangeL3NetworkDhcpIpAddress`, `ChangeL3NetworkState`, `CheckIpAvailability`, `GetAttachableVpcL3Network`, `GetCandidateL3NetworksForIpSecConnection`, `GetCandidateL3NetworksForLoadBalancer`, `GetCandidateL3NetworksForServerGroup`, `GetInterdependentL3NetworksBackupStorages`, `GetInterdependentL3NetworksImages`, `GetIpAddressCapacity`, `GetL3NetworkDhcpIpAddress`, `GetL3NetworkIpStatistic`, `GetL3NetworkMtu`, `GetL3NetworkRouterInterfaceIp`, `GetL3NetworkTypes`, `GetVmAttachableL3Network`, `IsVfNicAvailableInL3Network`, `RemoveDnsFromL3Network`, `RemoveHostRouteFromL3Network`, `SetL3NetworkMtu`, `SetL3NetworkRouterInterfaceIp` |
| [load_balancer_actions.go](../client/load_balancer_actions.go) | `AttachVipToLoadBalancer`, `GetLoadBalancerOwner`, `GetVpcAttachedLoadBalancer` |
| [load_balancer_listener_actions.go](../client/load_balancer_listener_actions.go) | `AddAccessControlListToLoadBalancer`, `AddCertificateToLoadBalancerListener`, `AddServerGroupToLoadBalancerListener`, `AddVmNicToLoadBalancer`, `GetCandidateVmNicsForLoadBalancer`, `GetLoadBalancerListenerACLEntries`, `RemoveAccessControlListFromLoadBalancer`, `RemoveCertificateFromLoadBalancerListener`, `RemoveServerGroupFromLoadBalancerListener`, `RemoveVmNicFromLoadBalancer` |
| [load_balancer_server_group_actions.go](../client/load_balancer_server_group_actions.go) | `AddBackendServerToServerGroup`, `ChangeLoadBalancerBackendServer`, `GetCandidateVmNicsForLoadBalancerServerGroup`, `RemoveBackendServerFromServerGroup` |
| [mdev_device_actions.go](../client/mdev_device_actions.go) | `AttachMdevDeviceToVm`, `DetachMdevDeviceFromVm`, `GenerateHygonMdevDevices`, `GenerateMdevDevices`, `GetMdevDeviceCandidates`, `UngenerateHygonMdevDevices`, `UngenerateMdevDevices` |
| [mdev_device_spec_actions.go](../client/mdev_device_spec_actions.go) | `AddMdevDeviceSpecToVmInstance`, `GetMdevDeviceSpecCandidates`, `RemoveMdevDeviceSpecFromVmInstance` |
| [multicast_router_actions.go](../client/multicast_router_actions.go) | `AddRendezvousPointToMulticastRouter`, `ChangeMulticastRouterState`, `GetVpcMulticastRoute`, `RemoveRendezvousPointFromMulticastRouter` |
| [network_diagnostics_actions.go](../client/network_diagnostics_actions.go) | `CheckNetworkReachable` |
| [network_service_l3network_ref_actions.go](../client/network_service_l3network_ref_actions.go) | `AttachNetworkServiceToL3Network`, `DetachNetworkServiceFromL3Network` |
| [network_service_provider_actions.go](../client/network_service_provider_actions.go) | `GetNetworkServiceTypes` |
| [ovn_controller_actions.go](../client/ovn_controller_actions.go) | `CreateOvnControllerVm` |
| [pci_device_actions.go](../client/pci_device_actions.go) | `AttachPciDeviceToVm`, `DetachPciDeviceFromVm`, `GenerateSriovPciDevices`, `GetEthernetVF`, `GetPciDeviceCandidatesForAttachingVm`, `GetPciDeviceCandidatesForNewCreateVm`, `PageEthernetVF`, `QueryEthernetVF`, `UngenerateSriovPciDevices` |
| [pci_device_pci_device_offering_ref_actions.go](../client/pci_device_pci_device_offering_ref_actions.go) | `GetPciDevicePciDeviceOffering`, `PagePciDevicePciDeviceOffering`, `QueryPciDevicePciDeviceOffering` |
| [pci_device_spec_actions.go](../client/pci_device_spec_actions.go) | `AddPciDeviceSpecToVmInstance`, `GetPciDeviceSpecCandidates`, `RemovePciDeviceSpecFromVmInstance` |
| [port_forwarding_rule_actions.go](../client/port_forwarding_rule_actions.go) | `ChangePortForwardingRuleState`, `GetVpcAttachedPortForwardingRules` |
| [port_mirror_actions.go](../client/port_mirror_actions.go) | `ChangePortMirrorState`, `GetCandidateVmNicsForPortMirror`, `GetPortMirrorNetworkUsedIp`, `PagePortMirrorNetworkUsedIp`, `QueryPortMirrorNetworkUsedIp` |
| [sdn_controller_actions.go](../client/sdn_controller_actions.go) | `PullHuaweiIMasterController`, `PullSdnControllerTenant`, `SdnControllerAddHost`, `SdnControllerChangeHost`, `SdnControllerRemoveHost` |
| [usb_device_actions.go](../client/usb_device_actions.go) | `AttachUsbDeviceToVm`, `DetachUsbDeviceFromVm`, `GetUsbDeviceCandidatesForAttachingVm` |
| [v2vconversion_host_actions.go](../client/v2vconversion_host_actions.go) | `ChangeV2VConversionHostState`, `CleanV2VConversionCache`, `ConvertVmFromForeignHypervisor`, `ConvertVmFromForeignHypervisorAsync` |
| [vip_actions.go](../client/vip_actions.go) | `ChangeVipState`, `CheckVipPortAvailability`, `GetVipAvailablePort`, `GetVipUsedPorts`, `GetVpcAttachedVip` |
| [virtual_router_vm_actions.go](../client/virtual_router_vm_actions.go) | `AddDnsToVpcRouter`, `CreateVpcVRouter`, `GetAttachablePublicL3ForVRouter`, `GetAvailableVpcL3Network`, `GetVpcRouter`, `GetVpcVRouterDistributedRoutingConnections`, `GetVpcVRouterDistributedRoutingEnabled`, `GetVpcVRouterNetworkServiceState`, `PageVpcRouter`, `ProvisionVirtualRouterConfig`, `QueryVpcRouter`, `ReconnectVirtualRouter`, `RemoveDnsFromVpcRouter`, `SetVpcVRouterDistributedRoutingEnabled`, `SetVpcVRouterNetworkServiceState`, `UpdateVirtualRouter` |
| [vm_cd_rom_actions.go](../client/vm_cd_rom_actions.go) | `AttachIsoToVmInstance`, `DetachIsoFromVmInstance`, `GetCandidateIsoForAttachingVm`, `GetCandidateVmForAttachingIso`, `SetVmInstanceDefaultCdRom` |
| [vm_instance_actions.go](../client/vm_instance_actions.go) | `AttachL3NetworkToVm`, `ChangeVmImage`, `ChangeVmPassword`, `CreateVmFromVmBackup`, `CreateVmFromVolumeBackup`, `CreateVmInstanceFromOvf`, `CreateVmInstanceFromOvfAsync`, `CreateVmInstanceFromVolume`, `CreateVmInstanceFromVolumeSnapshot`, `CreateVmInstanceFromVolumeSnapshotGroup`, `DeleteVmBootMode`, `DeleteVmConsolePassword`, `DeleteVmHostname`, `DeleteVmInstanceHaLevel`, `DeleteVmSshKey`, `DeleteVmStaticIp`, `DeleteVmUserDefinedXml`, `DescribeVmInstanceRecoveryPoint`, `ExecuteGuestVmCommand`, `FlattenVmInstance`, `FlattenVmInstanceAsync`, `FstrimVm`, `GetCandidateZonesClustersHostsForCreatingVm`, `GetVirtualizerInfo`, `GetVmBootOrder`, `GetVmCapabilities`, `GetVmConsolePassword`, `GetVmDeviceAddress`, `GetVmEmulatorPinning`, `GetVmHostname`, `GetVmInstanceFirstBootDevice`, `GetVmInstanceHaLevel`, `GetVmInstanceProtectedRecoveryPoints`, `GetVmInstanceRecoveryPoints`, `GetVmMonitorNumber`, `GetVmNuma`, `GetVmQga`, `GetVmRDP`, `GetVmSshKey`, `GetVmStartingCandidateClustersHosts`, `GetVmTask`, `GetVmUsbRedirect`, `GetVmXml`, `GetVmsCapabilities`, `GetVmvNUMATopology`, `ListVMsFromKVMHost`, `MigrateVm`, `MigrateVmAsync`, `MountVmInstanceRecoveryPoint`, `PauseVmInstance`, `ProtectVmInstanceRecoveryPoint`, `ReimageVmInstance`, `SetVmBootMode`, `SetVmBootOrder`, `SetVmBootVolume`, `SetVmCleanTraffic`, `SetVmClockTrack`, `SetVmConsoleMode`, `SetVmConsolePassword`, `SetVmEmulatorPinning`, `SetVmHostname`, `SetVmInstanceHaLevel`, `SetVmInstanceHygonMdev`, `SetVmMonitorNumber`, `SetVmNuma`, `SetVmQga`, `SetVmQxlMemory`, `SetVmRDP`, `SetVmSecurityLevel`, `SetVmSoundType`, `SetVmSshKey`, `SetVmStaticIp`, `SetVmUsbRedirect`, `SetVmUserDefinedXml`, `SyncVmClock`, `TakeVmConsoleScreenshot`, `UnmountVmInstanceRecoveryPoint`, `UnprotectVmInstanceRecoveryPoint`, `UpdateConsolePassword`, `UpdateVmNetworkConfig`, `UpdateVmPriority`, `UploadFileToVm` |
| [vm_nic_actions.go](../client/vm_nic_actions.go) | `AttachL3NetworkToVmNic`, `AttachVmNicToVm`, `ChangeVfNicHaState`, `ChangeVmNicNetwork`, `ChangeVmNicState`, `ChangeVmNicType`, `DeleteNicQos`, `DetachL3NetworkFromVm`, `GetCandidateL3NetworksForChangeVmNicNetwork`, `GetEipAttachableVmNics`, `GetNicQos`, `GetPortForwardingAttachableVmNics`, `GetVmNicAttachableEips`, `GetVmNicAttachedNetworkService`, `SetNicQos`, `UpdateVmNicDriver`, `UpdateVmNicMac` |
| [vm_priority_config_actions.go](../client/vm_priority_config_actions.go) | `UpdatePriorityConfig` |
| [vm_scheduling_rule_actions.go](../client/vm_scheduling_rule_actions.go) | `ChangeVmSchedulingRuleState`, `GetVmSchedulingRulesExecuteState`, `GetVmsSchedulingStateFromSchedulingRule`, `ListVmSchedulingRulesFromExecuteState`, `ListVmsFromSchedulingState` |
| [vm_scheduling_rule_group_actions.go](../client/vm_scheduling_rule_group_actions.go) | `AddVmToVmSchedulingRuleGroup`, `DetachVmFromVmSchedulingRuleGroup` |
| [vm_xml_hook_script_actions.go](../client/vm_xml_hook_script_actions.go) | `AttachUserDefinedXmlHookScriptToVm`, `CreateVmUserDefinedXmlHookScript`, `DeleteVmUserDefinedXmlHookScript`, `DetachUserDefinedXmlHookScriptFromVm`, `ExpungeVmUserDefinedXmlHookScript`, `GetVmUserDefinedXmlHookScript`, `GetVmXmlHookScript`, `PageVmUserDefinedXmlHookScript`, `QueryVmUserDefinedXmlHookScript`, `SetVmUserDefinedXmlHookScript`, `UpdateVmUserDefinedXmlHookScript` |
| [vpc_firewall_actions.go](../client/vpc_firewall_actions.go) | `DeleteFirewall`, `RefreshFirewall` |
| [vpc_firewall_ip_set_template_actions.go](../client/vpc_firewall_ip_set_template_actions.go) | `CreateFirewallIpSetTemplate`, `DeleteFirewallIpSetTemplate`, `GetFirewallIpSetTemplate`, `PageFirewallIpSetTemplate`, `QueryFirewallIpSetTemplate`, `UpdateFirewallIpSetTemplate` |
| [vpc_firewall_rule_actions.go](../client/vpc_firewall_rule_actions.go) | `ChangeFirewallRuleState`, `CheckFirewallRuleConfigFile`, `CreateFirewallRule`, `CreateFirewallRuleFromConfigFile`, `DeleteFirewallRule`, `GetFirewallRule`, `PageFirewallRule`, `QueryFirewallRule`, `UpdateFirewallRule` |
| [vpc_firewall_rule_set_actions.go](../client/vpc_firewall_rule_set_actions.go) | `ApplyRuleSetChanges`, `AttachFirewallRuleSetToL3`, `CreateFirewallRuleSet`, `DeleteFirewallRuleSet`, `DetachFirewallRuleSetFromL3`, `GetFirewallRuleSet`, `PageFirewallRuleSet`, `QueryFirewallRuleSet`, `UpdateFirewallRuleSet` |
| [vpc_firewall_rule_set_l3_ref_actions.go](../client/vpc_firewall_rule_set_l3_ref_actions.go) | `GetFirewallRuleSetL3Ref`, `PageFirewallRuleSetL3Ref`, `QueryFirewallRuleSetL3Ref` |
| [vpc_firewall_rule_template_actions.go](../client/vpc_firewall_rule_template_actions.go) | `CreateFirewallRuleTemplate`, `DeleteFirewallRuleTemplate`, `GetFirewallRuleTemplate`, `PageFirewallRuleTemplate`, `QueryFirewallRuleTemplate`, `UpdateFirewallRuleTemplate` |
| [vpc_ha_group_actions.go](../client/vpc_ha_group_actions.go) | `ChangeVpcHaGroupMonitorIps` |
| [vpc_shared_qos_actions.go](../client/vpc_shared_qos_actions.go) | `AttachVipToVpcSharedQos`, `ChangeVpcSharedQosBandwidth`, `DetachVipFromVpcSharedQos` |
| [vrouter_ospf_actions.go](../client/vrouter_ospf_actions.go) | `AddVRouterNetworksToOspfArea`, `CreateVRouterOspfArea`, `DeleteVRouterOspfArea`, `GetVRouterOspfArea`, `GetVRouterOspfNeighbor`, `GetVRouterOspfNetwork`, `GetVRouterRouterId`, `GetVpcAttachedOspf`, `PageVRouterOspfArea`, `PageVRouterOspfNetwork`, `QueryVRouterOspfArea`, `QueryVRouterOspfNetwork`, `RemoveVRouterNetworksFromOspfArea`, `SetVRouterRouterId`, `UpdateVRouterOspfArea` |
| [vrouter_route_table_actions.go](../client/vrouter_route_table_actions.go) | `AttachVRouterRouteTableToVRouter`, `DetachVRouterRouteTableFromVRouter`, `GetRouteTableVpcVRouterCandidate` |

## 存储与混合云

| 目标文件 | 方法 |
|---|---|
| [aliyun_connection_access_point_actions.go](../client/aliyun_connection_access_point_actions.go) | `AddConnectionAccessPointFromRemote`, `DeleteConnectionAccessPointLocal`, `GetConnectionAccessPointFromLocal`, `GetConnectionAccessPointFromRemote`, `PageConnectionAccessPointFromLocal`, `QueryConnectionAccessPointFromLocal`, `SyncConnectionAccessPointFromRemote` |
| [aliyun_disk_actions.go](../client/aliyun_disk_actions.go) | `AttachAliyunDiskToEcs`, `CreateAliyunDiskFromRemote`, `DeleteAliyunDiskFromLocal`, `DeleteAliyunDiskFromRemote`, `DetachAliyunDiskFromEcs`, `GetAliyunDiskFromLocal`, `PageAliyunDiskFromLocal`, `QueryAliyunDiskFromLocal`, `SyncDiskFromAliyunFromRemote` |
| [aliyun_nas_access_group_actions.go](../client/aliyun_nas_access_group_actions.go) | `GetAliyunNasAccessGroupRemote` |
| [aliyun_nas_access_group_rule_actions.go](../client/aliyun_nas_access_group_rule_actions.go) | `CreateAliyunNasAccessGroupRule`, `DeleteAliyunNasAccessGroupRule` |
| [aliyun_nas_file_system_actions.go](../client/aliyun_nas_file_system_actions.go) | `GetAliyunNasFileSystemRemote` |
| [aliyun_nas_mount_target_actions.go](../client/aliyun_nas_mount_target_actions.go) | `GetAliyunNasMountTargetRemote`, `UpdateAliyunMountTarget` |
| [aliyun_route_entry_actions.go](../client/aliyun_route_entry_actions.go) | `CreateAliyunVpcVirtualRouterEntryRemote`, `DeleteAliyunRouteEntryRemote`, `GetAliyunRouteEntryFromLocal`, `PageAliyunRouteEntryFromLocal`, `QueryAliyunRouteEntryFromLocal`, `SyncAliyunRouteEntryFromRemote` |
| [aliyun_router_interface_actions.go](../client/aliyun_router_interface_actions.go) | `CreateAliyunRouterInterfaceRemote`, `DeleteAliyunRouterInterfaceLocal`, `DeleteAliyunRouterInterfaceRemote`, `GetAliyunRouterInterfaceFromLocal`, `PageAliyunRouterInterfaceFromLocal`, `QueryAliyunRouterInterfaceFromLocal`, `StartConnectionBetweenAliyunRouterInterface`, `SyncAliyunRouterInterfaceFromRemote`, `UpdateAliyunRouteInterfaceRemote` |
| [aliyun_snapshot_actions.go](../client/aliyun_snapshot_actions.go) | `CreateAliyunSnapshotRemote`, `DeleteAliyunSnapshotFromLocal`, `DeleteAliyunSnapshotFromRemote`, `GCAliyunSnapshotRemote`, `GetAliyunSnapshotFromLocal`, `PageAliyunSnapshotFromLocal`, `QueryAliyunSnapshotFromLocal`, `SyncAliyunSnapshotRemote` |
| [aliyun_virtual_border_router_actions.go](../client/aliyun_virtual_border_router_actions.go) | `DeleteVirtualBorderRouterLocal`, `GetVirtualBorderRouterFromLocal`, `PageVirtualBorderRouterFromLocal`, `QueryVirtualBorderRouterFromLocal`, `RecoveryVirtualBorderRouterRemote`, `SyncVirtualBorderRouterFromRemote`, `TerminateVirtualBorderRouterRemote`, `UpdateVirtualBorderRouterRemote` |
| [aliyun_virtual_router_actions.go](../client/aliyun_virtual_router_actions.go) | `DeleteVirtualRouterLocal`, `GetAliyunVirtualRouterFromLocal`, `PageAliyunVirtualRouterFromLocal`, `QueryAliyunVirtualRouterFromLocal`, `SyncAliyunVirtualRouterFromRemote`, `UpdateAliyunVirtualRouter` |
| [aliyun_vswitch_connection_actions.go](../client/aliyun_vswitch_connection_actions.go) | `CreateConnectionBetweenL3NetworkAndAliyunVSwitch`, `DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch`, `GetConnectionBetweenL3NetworkAndAliyunVSwitch`, `PageConnectionBetweenL3NetworkAndAliyunVSwitch`, `QueryConnectionBetweenL3NetworkAndAliyunVSwitch`, `UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch` |
| [backup_storage_actions.go](../client/backup_storage_actions.go) | `AddSimulatorBackupStorage`, `AttachBackupStorageToZone`, `BackupStorageMigrateImage`, `BackupStorageMigrateImageAsync`, `ChangeBackupStorageState`, `CleanUpTrashOnBackupStorage`, `DeleteExportedImageFromBackupStorage`, `DetachBackupStorageFromZone`, `ExportImageFromBackupStorage`, `ExportImageFromBackupStorageAsync`, `GetBackupStorageCandidatesForImageMigration`, `GetBackupStorageCapacity`, `GetBackupStorageTypes`, `GetCandidateBackupStorageForCreatingImage`, `GetCdpBackupStorageRequirement`, `GetTrashOnBackupStorage`, `MergeDataOnBackupStorage` |
| [block_primary_storage_actions.go](../client/block_primary_storage_actions.go) | `AddStorageProtocol`, `GetBlockPrimaryStorageMetadata` |
| [block_volume_actions.go](../client/block_volume_actions.go) | `GetAccessPath` |
| [cdp_backup_actions.go](../client/cdp_backup_actions.go) | `CreateVmFromCdpBackup`, `CreateVmFromCdpBackupAsync`, `RevertVmFromCdpBackup`, `RevertVmFromCdpBackupAsync` |
| [cdp_task_actions.go](../client/cdp_task_actions.go) | `DeleteCdpTaskData`, `DisableCdpTask`, `EnableCdpTask`, `EnableCdpTaskAsync`, `UpgradeBackupStorageCdpTasks` |
| [ceph_backup_storage_actions.go](../client/ceph_backup_storage_actions.go) | `AddMonToCephBackupStorage`, `RemoveMonFromCephBackupStorage` |
| [ceph_primary_storage_actions.go](../client/ceph_primary_storage_actions.go) | `AddMonToCephPrimaryStorage`, `RemoveMonFromCephPrimaryStorage` |
| [container_image_actions.go](../client/container_image_actions.go) | `GetContainerUsage` |
| [data_integrity_actions.go](../client/data_integrity_actions.go) | `AddIntegrityResource`, `CheckBatchDataIntegrity` |
| [data_protection_actions.go](../client/data_protection_actions.go) | `StartDataProtection`, `StartDataProtectionAsync` |
| [database_backup_actions.go](../client/database_backup_actions.go) | `DeleteExportedDatabaseBackupFromBackupStorage`, `ExportDatabaseBackupFromBackupStorage`, `GetDatabaseBackupFromImageStore`, `RecoverDatabaseFromBackup`, `SyncDatabaseBackupFromImageStoreBackupStorage` |
| [ecs_image_actions.go](../client/ecs_image_actions.go) | `CreateEcsImageFromEcsSnapshot`, `CreateEcsImageFromLocalImage`, `DeleteEcsImageLocal`, `DeleteEcsImageRemote`, `GetCreateEcsImageProgress`, `GetEcsImageFromLocal`, `PageEcsImageFromLocal`, `QueryEcsImageFromLocal`, `SyncEcsImageFromRemote` |
| [ecs_instance_actions.go](../client/ecs_instance_actions.go) | `CreateEcsInstanceFromEcsImage`, `DeleteEcsInstanceLocal`, `GetEcsInstanceFromLocal`, `GetEcsInstanceType`, `GetEcsInstanceVncUrl`, `PageEcsInstanceFromLocal`, `QueryEcsInstanceFromLocal`, `SyncEcsInstanceFromRemote`, `UpdateEcsInstanceVncPassword` |
| [ecs_security_group_actions.go](../client/ecs_security_group_actions.go) | `CreateEcsSecurityGroupRemote`, `DeleteEcsSecurityGroupInLocal`, `DeleteEcsSecurityGroupRemote`, `GetEcsSecurityGroupFromLocal`, `PageEcsSecurityGroupFromLocal`, `QueryEcsSecurityGroupFromLocal`, `SyncEcsSecurityGroupFromRemote` |
| [ecs_security_group_rule_actions.go](../client/ecs_security_group_rule_actions.go) | `CreateEcsSecurityGroupRuleRemote`, `DeleteEcsSecurityGroupRuleRemote`, `GetEcsSecurityGroupRuleFromLocal`, `PageEcsSecurityGroupRuleFromLocal`, `QueryEcsSecurityGroupRuleFromLocal`, `SyncEcsSecurityGroupRuleFromRemote` |
| [ecs_vpc_actions.go](../client/ecs_vpc_actions.go) | `CreateEcsVpcRemote`, `DeleteEcsVpcInLocal`, `DeleteEcsVpcRemote`, `GetEcsVpcFromLocal`, `PageEcsVpcFromLocal`, `QueryEcsVpcFromLocal`, `SyncEcsVpcFromRemote` |
| [ecs_vswitch_actions.go](../client/ecs_vswitch_actions.go) | `CreateEcsVSwitchRemote`, `DeleteEcsVSwitchInLocal`, `DeleteEcsVSwitchRemote`, `GetEcsVSwitchFromLocal`, `PageEcsVSwitchFromLocal`, `QueryEcsVSwitchFromLocal`, `SyncEcsVSwitchFromRemote` |
| [external_primary_storage_actions.go](../client/external_primary_storage_actions.go) | `DiscoverExternalPrimaryStorage` |
| [hybrid_data_center_actions.go](../client/hybrid_data_center_actions.go) | `AddDataCenterFromRemote`, `DeleteAllEcsInstancesFromDataCenter`, `DeleteDataCenterInLocal`, `GetDataCenterFromLocal`, `GetDataCenterFromRemote`, `PageDataCenterFromLocal`, `QueryDataCenterFromLocal`, `SyncDataCenterFromRemote` |
| [hybrid_database_backup_actions.go](../client/hybrid_database_backup_actions.go) | `BackupDatabaseToPublicCloud`, `CreateOssBackupBucketRemote`, `DeleteBackupFileInPublic`, `DownloadBackupFileFromPublicCloud`, `GetOssBackupBucketFromRemote` |
| [hybrid_eip_actions.go](../client/hybrid_eip_actions.go) | `AttachHybridEipToEcs`, `CreateHybridEip`, `DeleteHybridEipFromLocal`, `DeleteHybridEipRemote`, `DetachHybridEipFromEcs`, `GetHybridEipFromLocal`, `PageHybridEipFromLocal`, `QueryHybridEipFromLocal`, `SyncHybridEipFromRemote`, `UpdateHybridEip` |
| [hybrid_identity_zone_actions.go](../client/hybrid_identity_zone_actions.go) | `AddIdentityZoneFromRemote`, `DeleteIdentityZoneInLocal`, `GetIdentityZoneFromLocal`, `GetIdentityZoneFromRemote`, `PageIdentityZoneFromLocal`, `QueryIdentityZoneFromLocal`, `SyncIdentityFromRemote` |
| [hybrid_key_actions.go](../client/hybrid_key_actions.go) | `AddAliyunKeySecret`, `AddHybridKeySecret`, `AttachAliyunKey`, `AttachHybridKey`, `DeleteAliyunKeySecret`, `DeleteHybridKeySecret`, `DetachAliyunKey`, `DetachHybridKey`, `GetHybridKeySecret`, `PageHybridKeySecret`, `QueryHybridKeySecret`, `UpdateAliyunKeySecret`, `UpdateHybridKeySecret` |
| [hybrid_vpc_vpn_connection_actions.go](../client/hybrid_vpc_vpn_connection_actions.go) | `CreateVpcVpnConnectionRemote`, `DeleteVpcVpnConnectionLocal`, `DeleteVpcVpnConnectionRemote`, `GetVpcVpnConfigurationFromRemote`, `GetVpcVpnConnectionFromLocal`, `PageVpcVpnConnectionFromLocal`, `QueryVpcVpnConnectionFromLocal`, `SyncVpcVpnConnectionFromRemote`, `UpdateVpcVpnConnectionRemote` |
| [hybrid_vpn_ike_config_actions.go](../client/hybrid_vpn_ike_config_actions.go) | `CreateVpnIkeConfig`, `DeleteVpcIkeConfigLocal`, `GetVpcIkeConfigFromLocal`, `PageVpcIkeConfigFromLocal`, `QueryVpcIkeConfigFromLocal` |
| [hybrid_vpn_ipsec_config_actions.go](../client/hybrid_vpn_ipsec_config_actions.go) | `CreateVpnIpsecConfig`, `DeleteVpcIpSecConfigLocal`, `GetVpcIpSecConfigFromLocal`, `PageVpcIpSecConfigFromLocal`, `QueryVpcIpSecConfigFromLocal` |
| [image_actions.go](../client/image_actions.go) | `CalculateImageHash`, `ChangeImageState`, `CreateDataVolumeTemplateFromVolume`, `CreateDataVolumeTemplateFromVolumeAsync`, `CreateDataVolumeTemplateFromVolumeBackup`, `CreateDataVolumeTemplateFromVolumeSnapshot`, `CreateRootVolumeTemplateFromRootVolume`, `CreateRootVolumeTemplateFromRootVolumeAsync`, `CreateRootVolumeTemplateFromVolumeBackup`, `CreateRootVolumeTemplateFromVolumeSnapshot`, `GetImageCandidatesForVmToChange`, `GetImageQga`, `GetUploadImageJobDetails`, `RecoveryImageFromImageStoreBackupStorage`, `SetImageBootMode`, `SetImageQga`, `SetImageSecurityLevel`, `SyncImageFromImageStoreBackupStorage`, `SyncImageSize` |
| [image_group_actions.go](../client/image_group_actions.go) | `CreateImageGroupFromImage`, `CreateImageGroupFromSnapshot`, `CreateImageGroupFromVmInstance` |
| [image_package_actions.go](../client/image_package_actions.go) | `ExportVmOvaPackage`, `ExportVmOvaPackageAsync`, `ParseOvf` |
| [image_replication_group_actions.go](../client/image_replication_group_actions.go) | `AddBackupStoragesToReplicationGroup` |
| [image_store_backup_storage_actions.go](../client/image_store_backup_storage_actions.go) | `AddDisasterImageStoreBackupStorage`, `GetImagesFromImageStoreBackupStorage`, `ReclaimSpaceFromImageStore`, `ReclaimSpaceFromImageStoreAsync`, `SetImageStoreBackupStorageQuota` |
| [iscsi_server_actions.go](../client/iscsi_server_actions.go) | `AttachIscsiServerToCluster`, `DetachIscsiServerFromCluster` |
| [local_raid_actions.go](../client/local_raid_actions.go) | `GetLocalRaidController`, `GetLocalRaidPhysicalDrive`, `GetLocalRaidPhysicalDriveSmart`, `GetPhysicalDriveSelfTestHistory`, `LocateLocalRaidPhysicalDrive`, `PageLocalRaidController`, `PageLocalRaidPhysicalDrive`, `PagePhysicalDriveSelfTestHistory`, `QueryLocalRaidController`, `QueryLocalRaidPhysicalDrive`, `QueryPhysicalDriveSelfTestHistory`, `RefreshLocalRaid`, `SelfTestLocalRaid` |
| [local_storage_resource_ref_actions.go](../client/local_storage_resource_ref_actions.go) | `GetLocalStorageHostDiskCapacity`, `LocalStorageGetVolumeMigratableHosts`, `LocalStorageMigrateVolume` |
| [mini_storage_actions.go](../client/mini_storage_actions.go) | `RecoverResourceSplitBrain` |
| [mini_storage_host_ref_actions.go](../client/mini_storage_host_ref_actions.go) | `BootstrapMiniHost`, `GetCandidateMiniHosts` |
| [nvme_server_actions.go](../client/nvme_server_actions.go) | `AttachNvmeServerToCluster`, `DetachNvmeServerFromCluster` |
| [oss_bucket_actions.go](../client/oss_bucket_actions.go) | `AddOssBucketFromRemote`, `AttachOssBucketToEcsDataCenter`, `CreateOssBucketRemote`, `DeleteOssBucketFileRemote`, `DeleteOssBucketNameLocal`, `DeleteOssBucketRemote`, `DetachOssBucketFromEcsDataCenter`, `GetOssBucketFileFromRemote`, `GetOssBucketFileName`, `GetOssBucketNameFromRemote`, `PageOssBucketFileName`, `QueryOssBucketFileName` |
| [primary_storage_actions.go](../client/primary_storage_actions.go) | `AddAliyunNasPrimaryStorage`, `AddLocalPrimaryStorage`, `AddNfsPrimaryStorage`, `AddSharedMountPointPrimaryStorage`, `AddSimulatorPrimaryStorage`, `AttachPrimaryStorageToCluster`, `ChangePrimaryStorageState`, `CleanUpImageCacheOnPrimaryStorage`, `CleanUpStorageTrashOnPrimaryStorage`, `CleanUpTrashOnPrimaryStorage`, `DetachPrimaryStorageFromCluster`, `GetCandidatePrimaryStoragesForCreatingVm`, `GetPrimaryStorageAllocatorStrategies`, `GetPrimaryStorageCandidatesForVmMigration`, `GetPrimaryStorageCandidatesForVolumeMigration`, `GetPrimaryStorageLicenseInfo`, `GetPrimaryStorageTypes`, `GetPrimaryStorageUsageReport`, `GetTrashOnPrimaryStorage`, `PrimaryStorageMigrateVm`, `PrimaryStorageMigrateVmAsync`, `PrimaryStorageMigrateVolume`, `PrimaryStorageMigrateVolumeAsync` |
| [scsi_lun_actions.go](../client/scsi_lun_actions.go) | `AttachScsiLunToVmInstance`, `CheckScsiLunClusterStatus`, `DetachScsiLunFromHost`, `DetachScsiLunFromVmInstance`, `GetScsiLunCandidatesForAttachingVm` |
| [shared_block_actions.go](../client/shared_block_actions.go) | `AddSharedBlockToSharedBlockGroup`, `GetSharedBlockCandidate`, `RefreshSharedblockDeviceCapacity` |
| [vm_backup_actions.go](../client/vm_backup_actions.go) | `CreateVmBackup`, `CreateVmBackupAsync`, `DeleteVmBackup`, `RecoverVmBackupFromImageStoreBackupStorage`, `RevertVmFromVmBackup`, `SyncVmBackup`, `SyncVmBackupFromImageStoreBackupStorage` |
| [volume_actions.go](../client/volume_actions.go) | `AttachDataVolumeToHost`, `AttachDataVolumeToVm`, `BatchSyncVolumeSize`, `ChangeVolumeState`, `CreateDataVolume`, `CreateDataVolumeFromVolumeBackup`, `CreateDataVolumeFromVolumeSnapshot`, `CreateDataVolumeFromVolumeTemplate`, `DeleteDataVolume`, `DeleteVolumeQos`, `DetachDataVolumeFromHost`, `DetachDataVolumeFromVm`, `ExportNbdVolumes`, `ExpungeDataVolume`, `FlattenVolume`, `FlattenVolumeAsync`, `GetDataVolumeAttachableVm`, `GetVmAttachableDataVolume`, `GetVolumeCapabilities`, `GetVolumeFormat`, `GetVolumeIoThreadPin`, `GetVolumeQos`, `RecoverDataVolume`, `ResizeDataVolume`, `ResizeRootVolume`, `SetVolumeIoThreadPin`, `SetVolumeQos`, `SyncVolumeSize`, `UnexportNbdVolumes` |
| [volume_backup_actions.go](../client/volume_backup_actions.go) | `RecoverBackupFromImageStoreBackupStorage`, `RevertVolumeFromVolumeBackup`, `SyncBackupFromImageStoreBackupStorage` |
| [volume_snapshot_actions.go](../client/volume_snapshot_actions.go) | `BatchDeleteVolumeSnapshot`, `CreateVolumesSnapshot`, `GetVolumeSnapshotSize`, `RevertVolumeFromSnapshot`, `RevertVolumeFromSnapshotAsync`, `ShrinkVolumeSnapshot`, `UndoSnapshotCreation`, `ValidateVolumeSnapshotChain` |
| [volume_snapshot_group_actions.go](../client/volume_snapshot_group_actions.go) | `CheckVolumeSnapshotGroupAvailability`, `GetMemorySnapshotGroupReference`, `RevertVmFromSnapshotGroup`, `UngroupVolumeSnapshotGroup` |
| [vpc_user_vpn_gateway_actions.go](../client/vpc_user_vpn_gateway_actions.go) | `CreateVpcUserVpnGatewayRemote`, `DeleteVpcUserVpnGatewayLocal`, `DeleteVpcUserVpnGatewayRemote`, `GetVpcUserVpnGatewayFromLocal`, `PageVpcUserVpnGatewayFromLocal`, `QueryVpcUserVpnGatewayFromLocal`, `SyncVpcUserVpnGatewayFromRemote` |
| [vpc_vpn_gateway_actions.go](../client/vpc_vpn_gateway_actions.go) | `DeleteVpcVpnGatewayLocal`, `GetVpcVpnGatewayFromLocal`, `PageVpcVpnGatewayFromLocal`, `QueryVpcVpnGatewayFromLocal`, `SyncVpcVpnGatewayFromRemote` |
| [zbox_actions.go](../client/zbox_actions.go) | `EjectZBox`, `SyncZBoxCapacity` |
| [zbox_backup_actions.go](../client/zbox_backup_actions.go) | `GetZBoxBackupDetails` |

## 身份权限与其他资源

| 目标文件 | 方法 |
|---|---|
| [access_control_list_actions.go](../client/access_control_list_actions.go) | `ChangeAccessControlListServerGroup` |
| [access_control_list_entry_actions.go](../client/access_control_list_entry_actions.go) | `AddAccessControlListRedirectRule`, `ChangeAccessControlListRedirectRule` |
| [access_key_actions.go](../client/access_key_actions.go) | `ChangeAccessKeyState` |
| [account_actions.go](../client/account_actions.go) | `CalculateAccountSpending`, `CheckApiPermission`, `CheckResourcePermission`, `CleanInvalidLdapBinding`, `GetAccountQuotaUsage`, `GetResourceAccount`, `LogOut`, `RefreshCaptcha`, `RenewSession`, `UnlockIdentity`, `ValidatePassword` |
| [account_resource_ref_actions.go](../client/account_resource_ref_actions.go) | `ChangeResourceOwner`, `RevokeResourceSharing`, `ShareResource` |
| [active_alarm_actions.go](../client/active_alarm_actions.go) | `ChangeActiveAlarmState`, `GetActiveAlarmStatus` |
| [affinity_group_actions.go](../client/affinity_group_actions.go) | `AddVmToAffinityGroup`, `ChangeAffinityGroupState`, `GetCandidateAffinityGroupForAttachingVm`, `GetCandidateAffinityGroupForCreatingVm`, `GetCandidateVMForAttachingAffinityGroup`, `RemoveVmFromAffinityGroup` |
| [alarm_actions.go](../client/alarm_actions.go) | `AddActionToAlarm`, `ChangeAlarmState`, `GetAlarmData`, `GetAlarmRecord`, `GetZWatchAlertHistogram`, `PageAlarmRecord`, `QueryAlarmRecord`, `RemoveActionFromAlarm`, `UpdateAlarmData` |
| [alarm_label_actions.go](../client/alarm_label_actions.go) | `AddLabelToAlarm`, `RemoveLabelFromAlarm` |
| [alert_data_ack_actions.go](../client/alert_data_ack_actions.go) | `AckAlarmData`, `AckEventData` |
| [app_build_system_actions.go](../client/app_build_system_actions.go) | `AttachAppBuildSystemToZone`, `ChangeAppBuildSystemState`, `DetachAppBuildSystemToZone`, `GetAppBuildSystemCapacity` |
| [application_development_service_actions.go](../client/application_development_service_actions.go) | `DeployAppDevelopmentService`, `SyncAINginxConfiguration` |
| [archive_ticket_actions.go](../client/archive_ticket_actions.go) | `GetArchiveTicketHistory`, `PageArchiveTicketHistory`, `QueryArchiveTicketHistory` |
| [audit_actions.go](../client/audit_actions.go) | `GetAudit`, `GetAuditData`, `PageAudit`, `QueryAudit` |
| [billing_actions.go](../client/billing_actions.go) | `CalculateAccountBillingSpending`, `CalculateResourceSpending`, `CleanupBillingUsage`, `GenerateAccountBilling`, `GetAccountBilling`, `PageAccountBilling`, `QueryAccountBilling` |
| [build_app_actions.go](../client/build_app_actions.go) | `AddBuildApp`, `CheckBuildAppParameters`, `CreateBuildApp`, `DeleteBuildApp`, `GetBuildApp`, `PageBuildApp`, `QueryBuildApp`, `UpdateBuildApp` |
| [build_app_export_history_actions.go](../client/build_app_export_history_actions.go) | `ExportBuildApp` |
| [cas_client_actions.go](../client/cas_client_actions.go) | `CreateOAuthClient`, `UpdateOAuthClient` |
| [cbt_task_actions.go](../client/cbt_task_actions.go) | `DisableCbtTask`, `EnableCbtTask` |
| [ccscertificate_actions.go](../client/ccscertificate_actions.go) | `AttachCCSCertificateToUser`, `DetachCCSCertificateFromUser`, `UpdateCCSCertificateUserState` |
| [certificate_actions.go](../client/certificate_actions.go) | `GetSpiceCertificates` |
| [chrony_actions.go](../client/chrony_actions.go) | `GetChronyServers`, `SyncChronyServers`, `UpdateChronyServers` |
| [console_proxy_agent_actions.go](../client/console_proxy_agent_actions.go) | `GetVmConsoleAddress`, `RequestConsoleAccess` |
| [container_management_endpoint_actions.go](../client/container_management_endpoint_actions.go) | `DeleteContainerResourceFromEndpoint` |
| [dataset_actions.go](../client/dataset_actions.go) | `DeleteDatasets`, `UpdateDatasets` |
| [debug_actions.go](../client/debug_actions.go) | `DebugSignal`, `GetDebugSignal` |
| [directory_actions.go](../client/directory_actions.go) | `AddResourcesToDirectory`, `MoveDirectory`, `MoveResourcesToDirectory`, `RemoveResourcesFromDirectory` |
| [disk_offering_actions.go](../client/disk_offering_actions.go) | `ChangeDiskOfferingState`, `ValidateDiskOfferingUserConfig` |
| [effective_resource_config_actions.go](../client/effective_resource_config_actions.go) | `GetResourceBindableConfig` |
| [encryption_actions.go](../client/encryption_actions.go) | `GetEncryptedField` |
| [error_elaboration_actions.go](../client/error_elaboration_actions.go) | `CheckElaborationContent`, `GetElaborationCategories`, `GetElaborations`, `ReloadElaboration` |
| [event_record_actions.go](../client/event_record_actions.go) | `GetEventRecord`, `PageEventRecord`, `QueryEventRecord` |
| [event_subscription_actions.go](../client/event_subscription_actions.go) | `AddActionToEventSubscription`, `ChangeEventSubscriptionState`, `GetAllEventMetadata`, `GetEventData`, `RemoveActionFromEventSubscription`, `RemoveLabelFromEventSubscription`, `SubscribeEvent`, `UnsubscribeEvent`, `UpdateEventData`, `UpdateSubscribeEvent` |
| [event_subscription_label_actions.go](../client/event_subscription_label_actions.go) | `AddLabelToEventSubscription` |
| [external_service_actions.go](../client/external_service_actions.go) | `GetExternalServices`, `ReloadExternalService` |
| [gc_job_actions.go](../client/gc_job_actions.go) | `DeleteGCJob`, `GetGCJob`, `PageGCJob`, `QueryGCJob`, `TriggerGCJob` |
| [global_config_actions.go](../client/global_config_actions.go) | `GetGlobalConfigOptions` |
| [guest_os_metadata_actions.go](../client/guest_os_metadata_actions.go) | `GetGuestOsMetadata`, `RefreshGuestOsMetadata` |
| [guest_tools_state_actions.go](../client/guest_tools_state_actions.go) | `AttachGuestToolsIsoToVm`, `GetLatestGuestToolsForVm`, `GetVmGuestToolsInfo` |
| [guest_vm_script_actions.go](../client/guest_vm_script_actions.go) | `ExecuteGuestVmScript`, `GetInvocationRecords` |
| [hygon_device_actions.go](../client/hygon_device_actions.go) | `GetHygonDevice`, `PageHygonDevice`, `QueryHygonDevice` |
| [iam2ldap_binding_actions.go](../client/iam2ldap_binding_actions.go) | `CreateIAM2VirtualIDLdapBinding`, `DeleteIAM2VirtualIDLdapBinding`, `GetCandidateLdapEntryForIAM2Binding`, `GetIAM2LdapBinding`, `PageIAM2LdapBinding`, `QueryIAM2LdapBinding` |
| [iam2organization_actions.go](../client/iam2organization_actions.go) | `ChangeIAM2OrganizationParent`, `ChangeIAM2OrganizationState`, `CheckIAM2OrganizationAvailability`, `GetIAM2OrganizationVirtualIDNumber`, `GetIAM2SystemAttributes`, `GetOrganizationQuotaUsage`, `SetOrganizationOperation`, `SetOrganizationSupervisor` |
| [iam2organization_attribute_actions.go](../client/iam2organization_attribute_actions.go) | `AddAttributesToIAM2Organization`, `RemoveAttributesFromIAM2Organization` |
| [iam2organization_project_ref_actions.go](../client/iam2organization_project_ref_actions.go) | `AttachIAM2ProjectToIAM2Organization`, `DetachIAM2ProjectFromIAM2Organization` |
| [iam2project_actions.go](../client/iam2project_actions.go) | `AddResourceToIAM2Project`, `ChangeIAM2ProjectState`, `CreateIAM2ProjectFromTemplate`, `GetIAM2ProjectContainerClusterCandidates`, `GetIAM2ProjectContainerImageTags`, `GetIAM2ProjectContainerImages`, `GetIAM2ProjectRepository`, `GetIAM2ProjectsOfVirtualID`, `SetIAM2ProjectContainerCluster`, `SetIAM2ProjectRetirePolicy`, `StopAllResourcesInIAM2Project` |
| [iam2project_attribute_actions.go](../client/iam2project_attribute_actions.go) | `AddAttributesToIAM2Project`, `RemoveAttributesFromIAM2Project` |
| [iam2project_template_actions.go](../client/iam2project_template_actions.go) | `CreateIAM2ProjectTemplateFromProject` |
| [iam2ticket_flow_collection_actions.go](../client/iam2ticket_flow_collection_actions.go) | `CreateIAM2TickFlowCollection` |
| [iam2virtual_id_actions.go](../client/iam2virtual_id_actions.go) | `AddIAM2VirtualIDsToOrganization`, `AddIAM2VirtualIDsToProject`, `AddIAM2VirtualIDsToProjects`, `AddRolesToIAM2VirtualID`, `BatchCreateIAM2VirtualIDFromConfigFile`, `ChangeIAM2VirtualIDState`, `ChangeIAM2VirtualIDType`, `CheckIAM2VirtualIDConfigFile`, `CleanInvalidLdapIAM2Binding`, `CreateIAM2VirtualIDFromLdapUid`, `GetIAM2VirtualIDAPIPermission`, `GetIAM2VirtualIDInGroup`, `RemoveIAM2VirtualIDsFromOrganization`, `RemoveIAM2VirtualIDsFromProject`, `RemoveIAM2VirtualIDsFromProjects`, `RemoveRolesFromIAM2VirtualID` |
| [iam2virtual_idattribute_actions.go](../client/iam2virtual_idattribute_actions.go) | `AddAttributesToIAM2VirtualID`, `RemoveAttributesFromIAM2VirtualID` |
| [iam2virtual_idgroup_actions.go](../client/iam2virtual_idgroup_actions.go) | `AddIAM2VirtualIDGroupToProjects`, `AddIAM2VirtualIDsToGroup`, `AddRolesToIAM2VirtualIDGroup`, `ChangeIAM2VirtualIDGroupState`, `RemoveIAM2VirtualIDGroupFromProjects`, `RemoveIAM2VirtualIDsFromGroup`, `RemoveRolesFromIAM2VirtualIDGroup` |
| [iam2virtual_idgroup_attribute_actions.go](../client/iam2virtual_idgroup_attribute_actions.go) | `AddAttributesToIAM2VirtualIDGroup`, `RemoveAttributesFromIAM2VirtualIDGroup` |
| [instance_offering_actions.go](../client/instance_offering_actions.go) | `ChangeInstanceOfferingState`, `CreateOvnControllerOffering`, `ValidateInstanceOfferingUserConfig` |
| [ldap_server_actions.go](../client/ldap_server_actions.go) | `CreateLdapBinding`, `DeleteLdapBinding`, `GetCandidateLdapEntryForBinding`, `GetLdapBinding`, `GetLdapEntry`, `GetLdapServerAvailableAttributes`, `PageLdapBinding`, `QueryLdapBinding` |
| [license_actions.go](../client/license_actions.go) | `DegradeFromLicenseServer`, `GetLicenseAddOns`, `GetLicenseCapabilities`, `GetLicenseInfo`, `GetLicenseNodeUsageDetails`, `GetLicenseRecords`, `GetLicenseUKeyStatus`, `IsLicenseServer`, `PushLicenseAddOnsUsage`, `RegisterLicenseRequestedApplication`, `RegisterLicenseServer`, `ReloadLicense`, `UnregisterLicenseRequestedApplication`, `UnregisterLicenseServer`, `VerifyLicenseServer`, `WithdrawLicenseCapacityApplication` |
| [license_authorized_capacity_actions.go](../client/license_authorized_capacity_actions.go) | `RequestLicenseCapacity`, `SyncLicenseCapacity` |
| [license_authorized_node_actions.go](../client/license_authorized_node_actions.go) | `UpgradeToLicenseServer` |
| [long_job_actions.go](../client/long_job_actions.go) | `CancelLongJob`, `RerunLongJob`, `RunIAM2Script`, `RunIAM2ScriptAsync`, `SubmitLongJob`, `SuspendLongJob` |
| [management_node_actions.go](../client/management_node_actions.go) | `GetCurrentTime`, `GetFactoryModeState`, `GetManagementNodeArch`, `GetManagementNodeDirCapacity`, `GetManagementNodeOS`, `GetPlatformTimeZone`, `GetSupportAPIs`, `GetVersion`, `IsOpensourceVersion`, `IsReadyToGo`, `UpdateFactoryModeState` |
| [media_actions.go](../client/media_actions.go) | `ChangeMediaState` |
| [metric_data_http_receiver_actions.go](../client/metric_data_http_receiver_actions.go) | `DeleteMetricData`, `PutMetricData` |
| [metric_query_actions.go](../client/metric_query_actions.go) | `GetAllMetricMetadata`, `GetMetricData`, `GetMetricLabelValue` |
| [model_actions.go](../client/model_actions.go) | `BindModelToService`, `DeleteModels`, `GenerateModelMetadata`, `GetSupportedIdentityModels`, `UnbindModelFromService` |
| [model_center_actions.go](../client/model_center_actions.go) | `GetModelCenterServices` |
| [model_eval_service_instance_group_actions.go](../client/model_eval_service_instance_group_actions.go) | `DeployModelEvalService` |
| [model_evaluation_task_actions.go](../client/model_evaluation_task_actions.go) | `DeleteModelEvaluationTasks` |
| [model_service_actions.go](../client/model_service_actions.go) | `DeleteModelServices`, `DeployDistributedModelService`, `DeployModelService`, `GetMaaSUsage`, `MatchModelServiceTemplateWithModel`, `RestartModelServiceGroups` |
| [model_service_instance_group_actions.go](../client/model_service_instance_group_actions.go) | `DeleteModelServiceInstanceGroups` |
| [monitor_group_instance_actions.go](../client/monitor_group_instance_actions.go) | `AddInstanceToMonitorGroup`, `RemoveInstanceFromMonitorGroup` |
| [monitor_template_actions.go](../client/monitor_template_actions.go) | `ApplyMonitorTemplateToMonitorGroup`, `RevokeMonitorTemplateFromMonitorGroup` |
| [monitor_trigger_actions.go](../client/monitor_trigger_actions.go) | `AttachMonitorTriggerActionToTrigger`, `ChangeMonitorTriggerState`, `DetachMonitorTriggerActionFromTrigger`, `GetMonitorItem` |
| [monitor_trigger_action_actions.go](../client/monitor_trigger_action_actions.go) | `ChangeMonitorTriggerActionState`, `CreateEmailMonitorTriggerAction`, `GetEmailTriggerAction`, `PageEmailTriggerAction`, `QueryEmailTriggerAction`, `UpdateEmailMonitorTriggerAction` |
| [mtty_device_actions.go](../client/mtty_device_actions.go) | `GenerateSeMdevDevices`, `UngenerateSeMdevDevices` |
| [nfv_inst_actions.go](../client/nfv_inst_actions.go) | `ProvisionNfvInstConfig`, `UpdateNfvInstProvisionConfig` |
| [nfv_inst_group_actions.go](../client/nfv_inst_group_actions.go) | `AttachNfvInstToGroup`, `ChangeNfvInstGroupOperationMode`, `DetachNfvInstFromGroup`, `ProvisionNfvInstGroup` |
| [oauth2token_actions.go](../client/oauth2token_actions.go) | `TokenIntrospection` |
| [observability_server_actions.go](../client/observability_server_actions.go) | `AttachServiceToObservabilityServer`, `CreateObservabilityServer`, `DetachServiceFromObservabilityServer`, `GetObservabilityServerServiceData` |
| [plugin_driver_actions.go](../client/plugin_driver_actions.go) | `DeletePluginDrivers`, `GetPluginDrivers`, `PagePluginDrivers`, `QueryPluginDrivers`, `RefreshPluginDrivers` |
| [policy_route_rule_set_actions.go](../client/policy_route_rule_set_actions.go) | `AttachPolicyRouteRuleSetToL3`, `DetachPolicyRouteRuleSetFromL3`, `GetPolicyRouteRuleSetFromVirtualRouter` |
| [preconfiguration_template_actions.go](../client/preconfiguration_template_actions.go) | `ChangePreconfigurationTemplateState` |
| [price_table_actions.go](../client/price_table_actions.go) | `AttachPriceTableToAccount`, `ChangeAccountPriceTableBinding`, `DetachPriceTableFromAccount`, `ValidatePriceUserConfig` |
| [prometheus_actions.go](../client/prometheus_actions.go) | `GetPrometheusMetricLabelValue`, `PrometheusQueryLabelValues`, `PrometheusQueryMetadata`, `PrometheusQueryPassThrough`, `PrometheusQueryVmMonitoringData` |
| [publish_app_actions.go](../client/publish_app_actions.go) | `GetResourceFromPublishApp`, `PreviewResourceFromApp` |
| [query_actions.go](../client/query_actions.go) | `BatchQuery`, `ZQLQuery` |
| [quota_actions.go](../client/quota_actions.go) | `UpdateOrganizationQuota` |
| [resource_actions.go](../client/resource_actions.go) | `GetResourceNames` |
| [resource_config_actions.go](../client/resource_config_actions.go) | `GetResourceConfigs`, `UpdateResourceConfigs` |
| [resource_price_actions.go](../client/resource_price_actions.go) | `CreateResourcePrice`, `DeleteResourcePrice`, `GetResourcePrice`, `PageResourcePrice`, `QueryResourcePrice`, `UpdateResourcePrice` |
| [resource_stack_actions.go](../client/resource_stack_actions.go) | `AddResourceStackVmPortMonitor`, `CreateResourceStackFromApp`, `DeleteResourceStackVmPortMonitor`, `GetEventFromResourceStack`, `GetResourceFromResourceStack`, `GetResourceStackFromResource`, `GetResourceStackVmStatus`, `GetSupportedCloudFormationResources`, `PageEventFromResourceStack`, `PreviewResourceStack`, `QueryEventFromResourceStack`, `RestartResourceStack` |
| [role_actions.go](../client/role_actions.go) | `AddPolicyStatementsToRole`, `AttachPolicyToRole`, `AttachRoleToAccount`, `ChangeRoleState`, `DetachPolicyFromRole`, `DetachRoleFromAccount`, `RemovePolicyStatementsFromRole` |
| [scheduler_job_actions.go](../client/scheduler_job_actions.go) | `AddSchedulerJobToSchedulerTrigger`, `ChangeSchedulerState`, `GetNoTriggerSchedulerJobs`, `RemoveSchedulerJobFromSchedulerTrigger` |
| [scheduler_job_group_actions.go](../client/scheduler_job_group_actions.go) | `AddSchedulerJobGroupToSchedulerTrigger`, `AddSchedulerJobsToSchedulerJobGroup`, `RemoveSchedulerJobGroupFromSchedulerTrigger`, `RemoveSchedulerJobsFromSchedulerJobGroup` |
| [scheduler_job_history_actions.go](../client/scheduler_job_history_actions.go) | `GetSchedulerExecutionReport` |
| [scheduler_trigger_actions.go](../client/scheduler_trigger_actions.go) | `GetAvailableTriggers`, `RunSchedulerTrigger` |
| [search_index_actions.go](../client/search_index_actions.go) | `RefreshSearchIndexes` |
| [secret_resource_pool_actions.go](../client/secret_resource_pool_actions.go) | `ChangeSecretResourcePoolState`, `CreateJitSecretResourcePool`, `GetSignatureServerEncryptPublicKey` |
| [security_group_actions.go](../client/security_group_actions.go) | `AddVmNicToSecurityGroup`, `AttachSecurityGroupToL3Network`, `ChangeSecurityGroupState`, `DeleteVmNicFromSecurityGroup`, `DetachSecurityGroupFromL3Network`, `GetCandidateVmNicForSecurityGroup`, `GetVmNicInSecurityGroup`, `PageVmNicInSecurityGroup`, `QueryVmNicInSecurityGroup`, `SetVmNicSecurityGroup` |
| [security_group_rule_actions.go](../client/security_group_rule_actions.go) | `ChangeSecurityGroupRuleState`, `UpdateSecurityGroupRulePriority` |
| [security_machine_actions.go](../client/security_machine_actions.go) | `ChangeSecurityMachineState`, `SecurityMachineDetectSync`, `SecurityMachineEncrypt`, `SetSecurityMachineKey` |
| [slb_group_actions.go](../client/slb_group_actions.go) | `ChangeSlbGroupDeployType`, `ChangeSlbGroupMonitorIps` |
| [slb_vm_instance_actions.go](../client/slb_vm_instance_actions.go) | `CreateSlbInstance`, `ProvisionSlbInstance` |
| [snsaliyun_sms_endpoint_actions.go](../client/snsaliyun_sms_endpoint_actions.go) | `GetSNSSmsEndpoint`, `PageSNSSmsEndpoint`, `QuerySNSSmsEndpoint` |
| [snsapplication_endpoint_actions.go](../client/snsapplication_endpoint_actions.go) | `ChangeSNSApplicationEndpointState` |
| [snsapplication_platform_actions.go](../client/snsapplication_platform_actions.go) | `ChangeSNSApplicationPlatformState` |
| [snsding_talk_at_person_actions.go](../client/snsding_talk_at_person_actions.go) | `UpdateAtPersonOfAtDingTalkEndpoint` |
| [snsding_talk_endpoint_actions.go](../client/snsding_talk_endpoint_actions.go) | `SNSDingTalkTestConnection` |
| [snsemail_address_actions.go](../client/snsemail_address_actions.go) | `AddEmailAddressToSNSEmailEndpoint`, `DeleteEmailAddressOfSNSEmailEndpoint`, `UpdateEmailAddressOfSNSEmailEndpoint` |
| [snsemail_endpoint_actions.go](../client/snsemail_endpoint_actions.go) | `SNSEmailTestConnection` |
| [snsendpoint_thirdparty_alert_history_actions.go](../client/snsendpoint_thirdparty_alert_history_actions.go) | `GetThirdpartyAlert`, `PageThirdpartyAlert`, `QueryThirdpartyAlert`, `UpdateThirdpartyAlerts` |
| [snsfei_shu_at_person_actions.go](../client/snsfei_shu_at_person_actions.go) | `UpdateAtPersonOfAtFeiShuEndpoint` |
| [snsfei_shu_endpoint_actions.go](../client/snsfei_shu_endpoint_actions.go) | `SNSFeiShuTestConnection` |
| [snshttp_endpoint_actions.go](../client/snshttp_endpoint_actions.go) | `SNSHttpTestConnection` |
| [snsmicrosoft_teams_endpoint_actions.go](../client/snsmicrosoft_teams_endpoint_actions.go) | `SNSMicrosoftTeamsTestConnection` |
| [snssnmp_endpoint_actions.go](../client/snssnmp_endpoint_actions.go) | `CreateSNSSnmpEndpoint`, `SNSSnmpTestConnection` |
| [snstext_template_actions.go](../client/snstext_template_actions.go) | `GetTextTemplateArg` |
| [snstopic_actions.go](../client/snstopic_actions.go) | `ChangeSNSTopicState`, `GetSNSTopicSubscriber`, `PageSNSTopicSubscriber`, `QuerySNSTopicSubscriber`, `SubscribeSNSTopic`, `UnsubscribeSNSTopic` |
| [snswe_com_at_person_actions.go](../client/snswe_com_at_person_actions.go) | `UpdateAtPersonOfAtWeComEndpoint` |
| [snswe_com_endpoint_actions.go](../client/snswe_com_endpoint_actions.go) | `SNSWeComTestConnection` |
| [ssh_key_pair_actions.go](../client/ssh_key_pair_actions.go) | `AttachSshKeyPairToVmInstance`, `DetachSshKeyPairFromVmInstance` |
| [ssoclient_actions.go](../client/ssoclient_actions.go) | `SsoClientPushData` |
| [stack_template_actions.go](../client/stack_template_actions.go) | `CheckStackTemplateParameters`, `DecodeStackTemplate` |
| [system_tag_actions.go](../client/system_tag_actions.go) | `CreateSystemTags` |
| [tag_actions.go](../client/tag_actions.go) | `AttachTagToResources`, `CreateTag`, `DeleteTag`, `DetachTagFromResources`, `GetTag`, `PageTag`, `QueryTag`, `UpdateTag` |
| [task_actions.go](../client/task_actions.go) | `CleanQueue`, `GetChainTask`, `GetTaskProgress` |
| [ticket_actions.go](../client/ticket_actions.go) | `ChangeTicketStatus`, `GetTicketHistory`, `PageTicketHistory`, `QueryTicketHistory`, `UpdateTicketRequest` |
| [ticket_flow_collection_actions.go](../client/ticket_flow_collection_actions.go) | `AddTicketTypesToTicketFlowCollection`, `ChangeTicketFlowCollectionState`, `RemoveTicketTypesFromTicketFlowCollection` |
| [two_factor_authentication_actions.go](../client/two_factor_authentication_actions.go) | `GetTwoFactorAuthenticationState` |
| [user_actions.go](../client/user_actions.go) | `AttachPoliciesToUser`, `AttachPolicyToUser`, `DetachPoliciesFromUser`, `DetachPolicyFromUser` |
| [user_group_actions.go](../client/user_group_actions.go) | `AddUserToGroup`, `AttachPolicyToUserGroup`, `DetachPolicyFromUserGroup`, `RemoveUserFromGroup` |
| [user_proxy_config_actions.go](../client/user_proxy_config_actions.go) | `AddProxyToResource` |
| [vtep_actions.go](../client/vtep_actions.go) | `CreateVxlanVtep` |
| [zone_actions.go](../client/zone_actions.go) | `ChangeZoneState` |
