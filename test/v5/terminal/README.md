# Terminal V5 测试用例

本目录包含了终端设备相关接口的测试用例。

## 测试覆盖范围

### 设备管理接口
- `TestList` - 获取设备列表
- `TestDeviceJournalStat` - 设备日志统计

### 设备统计接口
- `TestDeviceSummaryByCategory` - 按设备分类统计
- `TestDeviceSummaryByNetStatus` - 按网络状态统计设备

## 运行测试

### 运行所有测试
```bash
go test -v ./test/v5/terminal/
```

### 运行单个测试
```bash
go test -v ./test/v5/terminal/ -run TestList
```

## 配置说明

在运行测试前，请确保在 `terminal_test.go` 中配置正确的测试参数：

- `accessId` - 访问ID
- `accessKey` - 访问密钥
- `token` - 认证令牌
- `sid` - 景区ID列表

## 注意事项

1. 测试用例使用的是模拟数据，实际运行时请根据真实环境调整参数
2. 确保测试环境网络连接正常，能够访问 ODAS API 服务
3. 设备相关接口需要确保有相应的设备数据