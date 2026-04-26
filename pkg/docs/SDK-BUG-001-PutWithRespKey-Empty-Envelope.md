# SDK-BUG-001 — `PutWithRespKey` / `PutWithSpec` 空 `responseKey` 导致 inventory envelope 解析失败

| 字段 | 值 |
|---|---|
| 严重度 | **High** |
| 状态 | ✅ 已修（2026-04-26，本仓库当前工作树） |
| 发现日期 | 2026-04-15（`zstack_alarm` Update 测试） |
| 影响 SDK 文件 | `pkg/client/http_client.go` + 569 处 `*_actions.go` 调用 |
| 影响 provider 资源数 | ≥ 23（已绕过） |
| 跟踪关键字 | `SDK-BUG-001`、`SDK-WA-001`、`PutWithRespKey.*""` |

---

## 1. 现象

任何使用 `cli.PutWithRespKey(path, uuid, "", params, &resp)` 的 SDK 方法，在调用后返回的 `resp` 为**全零值 struct**（所有字段空字符串 / 0 / false），即使 ZStack API 实际返回了完整的 inventory 数据。

下游 provider 把空 struct 写回 Terraform state，会触发：

```
Error: Provider produced inconsistent result after apply
  .uuid:  was "688a692a...", but now ""
  .name:  was "acc-test-alarm", but now ""
  ...
```

---

## 2. 根因

### 2.1 ZStack API 响应格式

ZStack PUT 类 API 的标准响应是带 `inventory` envelope 的 JSON：

```json
{
  "inventory": {
    "uuid": "...",
    "name": "...",
    "...": "..."
  }
}
```

### 2.2 SDK HTTP 层解析逻辑

`pkg/client/http_client.go` 第 360–377 行 `PutWithAsync`：

```go
func (cli *ZSHttpClient) PutWithAsync(ctx context.Context, resource, resourceId, spec, responseKey string, params interface{}, retVal interface{}, async bool) (string, error) {
    urlStr := cli.getPutURL(resource, resourceId, spec)
    location, _, resp, err := cli.httpPut(ctx, urlStr, jsonMarshal(params), async)
    if err != nil {
        return location, err
    }

    if async || retVal == nil {
        return location, nil
    }

    if len(responseKey) == 0 {
        return location, resp.Unmarshal(retVal)        // ← 缺 envelope 提取
    }

    return location, resp.Unmarshal(retVal, responseKey) // ← 正确路径
}
```

当 `responseKey == ""` 时，SDK 直接把 `{"inventory": {...}}` 整体 unmarshal 到目标 view struct（例如 `view.AlarmInventoryView`）。该 struct 没有 `inventory` 字段，匹配失败，全部字段为零值。

### 2.3 大量 `*_actions.go` 都传了空 responseKey

```bash
$ grep -rn 'PutWithRespKey.*""' pkg/client/ | wc -l
569
```

例如 `alarm_actions.go`：

```go
func (cli *ZSClient) UpdateAlarm(uuid string, params param.UpdateAlarmParam) (*view.AlarmInventoryView, error) {
    resp := view.AlarmInventoryView{}
    if err := cli.PutWithRespKey("v1/zwatch/alarms", uuid, "", map[string]interface{}{
        //                                            ↑ 应为 "inventory"
        "updateAlarm": params.Params,
    }, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}
```

`Get*` / `List*` / `Page*` 类方法没有这个问题，因为它们走 `responseKeyInventory` 常量；只有自动生成的 `Update*` / 部分 action 类 `Put*` 没传。

---

## 3. 复现

```bash
# 在 provider 仓库
cd terraform-provider-zstack
git checkout fix/qa-20260422-p0-plus-story15  # workaround 之前的最初提交
TF_ACC=1 go test -tags=integration ./zstack/provider/ -run TestAccAlarmResource -v
# Step 2 (Update) 必失败：state diff 全部归零
```

---

## 4. 修复方案

### 方案 A — 一行式

每个 `Update*` / 受影响的 `Put*` action 把 `""` 改成 `"inventory"`：

```diff
 func (cli *ZSClient) UpdateAlarm(uuid string, params param.UpdateAlarmParam) (*view.AlarmInventoryView, error) {
     resp := view.AlarmInventoryView{}
-    if err := cli.PutWithRespKey("v1/zwatch/alarms", uuid, "", map[string]interface{}{
+    if err := cli.PutWithRespKey("v1/zwatch/alarms", uuid, "inventory", map[string]interface{}{
         "updateAlarm": params.Params,
     }, &resp); err != nil {
         return nil, err
     }
     return &resp, nil
 }
```

需要扫描全部 569 处，逐个判断响应是否真的有 `inventory` envelope。建议给 SDK 代码生成器加一条规则：默认 `Put*` action 用 `"inventory"`。

### 方案 B — HTTP 层兜底（已采用）

修改 `PutWithAsync` 在 `responseKey == ""` 时尝试默认 `"inventory"`；找不到再回退到当前的整体 unmarshal：

```go
if len(responseKey) == 0 {
    // 尝试 inventory envelope
    if resp.Contains("inventory") {
        return location, resp.Unmarshal(retVal, "inventory")
    }
    return location, resp.Unmarshal(retVal)
}
```

优点：不用动 569 处 action；缺点：埋了一层隐式行为，将来 ZStack 其他 envelope 名出现时还会出问题。

当前仓库已在 `pkg/client/http_client.go` 的 `PutWithAsync` 落地兜底：当 `responseKey == ""` 且响应体包含 `inventory` 时，优先提取 `inventory` 再反序列化；否则维持原有整体 unmarshal 行为。

> 后续仍建议把代码生成器规则补齐，避免继续生成空 `responseKey` 的 `Put*` action。

---

## 5. 已确认受影响的 provider 资源（23 处）

> 这些资源的 Update（部分含 Create）方法已经在 provider 侧追加 `findResourceByGet` / `findResourceByQuery` 重读 inventory 来绕过此 SDK bug。

| Provider 文件 | SDK 方法 | Provider workaround |
|---|---|---|
| `resource_zstack_account.go` | `UpdateAccount` | `_, err := Update(...)` → `GetAccount(uuid)` |
| `resource_zstack_affinity_group.go` | `UpdateAffinityGroup` | `_, err :=` → `GetAffinityGroup` |
| `resource_zstack_alarm.go` | `UpdateAlarm` | `_, err :=` → `findResourceByQuery(QueryAlarm)` |
| `resource_zstack_auto_scaling_group.go` | `UpdateAutoScalingGroup` | `_, err :=` → `GetAutoScalingGroup` |
| `resource_zstack_backup_storage.go` | `UpdateBackupStorage` | `_, err :=` |
| `resource_zstack_cluster.go` | `UpdateCluster` | `_, err :=` |
| `resource_zstack_global_config.go` | `UpdateGlobalConfig` (Create+Update) | `_, err :=` → `findResourceByQuery(QueryGlobalConfig)` |
| `resource_zstack_host.go` | `UpdateHost` / `UpdateKVMHost` | `_, err :=` |
| `resource_zstack_iam2_project.go` | `UpdateIAM2Project` | `_, err :=` |
| `resource_zstack_image_store_backup_storage.go` | `UpdateImageStoreBackupStorage` | `_, err :=` |
| `resource_zstack_instance.go` | `UpdateVmInstance` | `_, err :=` → `findResourceByGet(GetVmInstance)` |
| `resource_zstack_instance_scripts.go` | `UpdateGuestVmScript` | `_, err :=` |
| `resource_zstack_l2vlan_network.go` | `UpdateL2Network` | `_, err :=` |
| `resource_zstack_l3network.go` | `UpdateL3Network` | `_, err :=` → `findResourceByQuery(QueryL3Network)` |
| `resource_zstack_load_balancer.go` | `UpdateLoadBalancer` | `_, err :=` → `GetLoadBalancer` |
| `resource_zstack_load_balancer_listener.go` | `UpdateLoadBalancerListener` | `_, err :=` → `GetLoadBalancerListener` |
| `resource_zstack_port_forwarding_rule.go` | `UpdatePortForwardingRule` | `_, err :=` |
| `resource_zstack_primary_storage.go` | `UpdatePrimaryStorage` | `_, err :=` |
| `resource_zstack_ssh_key_pair.go` | `UpdateSshKeyPair` | `_, err :=` |
| `resource_zstack_vm_cdrom.go` | `UpdateVmCdRom` | `_, err :=` → `findResourceByQuery(QueryVmCdRom)` |
| `resource_zstack_volume.go` | `UpdateVolume` | `_, err :=` |
| `resource_zstack_volume_snapshot.go` | `UpdateVolumeSnapshot` | `_, err :=` |
| `resource_zstack_zone.go` | `UpdateZone` | `_, err :=` |

> 完整列表见 `terraform-provider-zstack/_bmad-output/bug-tracker.md` § "SDK-WA-001"。

---

## 6. SDK 修复后的 Provider 回收清单

### 当前状态

- SDK 侧已修复 `PutWithRespKey(..., "", ...)` 的 `inventory` envelope 回退解析。
- 本仓库已补回归测试：`pkg/client/client_behavior_test.go`
- 下游 provider 可开始逐步移除“Update 后再 `Get*` / `Query*` 重读”的临时 workaround。

SDK 修复后，下游 provider 可：

```bash
cd terraform-provider-zstack
# 1. 升级 SDK
go get github.com/zstackio/zstack-sdk-go-v2@latest

# 2. 找回收点（Update 后的 re-query workaround）
grep -rn "_, err := r.client.Update" zstack/provider/

# 3. 还原模式
# 把:
#   _, err := r.client.UpdateXxx(uuid, params)
#   ...err handle...
#   inv, err := findResourceByQuery(r.client.QueryXxx, uuid)
#
# 改回:
#   inv, err := r.client.UpdateXxx(uuid, params)

# 4. 删 troubleshooting/SDK-BUG-UpdateAlarm-Empty-Response.md 的方案 B 段落
```

或者**保守做法**：SDK 修了之后保留 `findResourceByQuery` 当成"刷新最新状态"的一致性保障，只删除 `// SDK bug:` 注释 — 因为多一次 query 的网络成本可忽略，但 state 一致性更稳。

---

## 7. 原始报告

- `terraform-provider-zstack/troubleshooting/SDK-BUG-UpdateAlarm-Empty-Response.md` — `zstack_alarm` 案例完整调查日志
- `terraform-provider-zstack/_bmad-output/bug-tracker.md` § BUG-056 / BUG-061 — 影响范围扫描记录
