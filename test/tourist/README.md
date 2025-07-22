# Tourist V5 测试用例

本目录包含了游客分析相关接口的测试用例。

## 测试覆盖范围

### 游客预测接口
- `TestForecastTouristSummary` - 游客预测汇总
- `TestPreBookingSummary` - 预订汇总统计
- `TestLastDayTemporaryBookingSummary` - 昨日临时预订汇总

### 游客地域分析接口
- `TestInsideAndOutsideByProvince` - 按省份内外游客统计
- `TestTouristSummaryByProvince` - 按省份游客汇总
- `TestTouristSummaryByCity` - 按城市游客汇总
- `TestTouristSummaryByDistrict` - 按区县游客汇总
- `TestTouristLocalByTicket` - 按票务本地游客统计

### 游客属性分析接口
- `TestSexSummaryByAge` - 按年龄性别汇总
- `TestSummaryByType` - 按类型汇总游客
- `TestTouristSummaryByPeer` - 按同行人数汇总游客

### 游客消费分析接口
- `TestTicketSummaryByPayChannel` - 按支付渠道票务汇总
- `TestTouristSummary` - 游客汇总统计

## 运行测试

### 运行所有测试
```bash
go test -v ./test/v5/tourist/
```

### 运行单个测试
```bash
go test -v ./test/v5/tourist/ -run TestTouristSummary
```

## 配置说明

在运行测试前，请确保在 `tourist_test.go` 中配置正确的测试参数：

- `accessId` - 访问ID
- `accessKey` - 访问密钥
- `token` - 认证令牌
- `sid` - 景区ID列表

## 注意事项

1. 测试用例使用的是模拟数据，实际运行时请根据真实环境调整参数
2. 确保测试环境网络连接正常，能够访问 ODAS API 服务
3. 游客分析接口需要有足够的历史数据支撑
4. 地域分析功能依赖于游客身份证或手机号归属地数据