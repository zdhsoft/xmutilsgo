# xmutilsgo

## 使用办法

```bash
go get github.com/zdhsoft/xmutilsgo
```

## 主要内容

- int.go 定义泛型的整数类型和字符串转整数的函数和随机范围的函数
- isin.go 判断指定元素是否再数组中的函数
- page.go mysql 用于分页的类
- ret.go 通用返回值的类
- set.go 基于 map 实现的集合功能
- string.go 主要是字符串链接和 pad 的函数
- time.go 用于时间处理的功能函数
- array.go 数组相关工具函数
- map.go map 相关工具函数
- json.go json 相关工具函数
- md5.go md5 相关工具函数
- utils.go 其他工具函数
- version.go 版本信息
- request.go 封装了一些常用的 http 请求函数
- gorm.go 封装了一些常用的 gorm 操作函数
- cmp.go 比较相关的函数

## 版本信息

- v.1.0.10 增加雪花算法
  - 雪花算法实现：
    - 使用64位整数存储ID
    - 包含时间戳（41位）、机器ID（10位）和序列号（12位）
    - 支持分布式环境下的唯一性
    - 时间戳从2024年开始，可以使用约69年
  - 分表策略：
    - 使用 GetShardKey 函数根据ID计算分表键
    - 采用取模运算，确保数据均匀分布
    -分表数量建议为2的幂次方（如16、32、64等）

- v1.0.9 增加比较相关的函数
  - Cmp 比较类的模板函数，支持所有整数，浮点数和字符串
    - CmpTime 比较时间的函数
    - CmpBool 比较 bool 型的函数

- v1.0.8 增加了 gorm 的方法

  - AddDateScopeDateTime 增加日期范围条件，被查询的字段是 datetime 类型
  - 注释增加了查询字段的类型说明

- v1.0.7 增加一些 0 点计算的函数

  - 增加 GetMidnightTimeToTime 和 GetMidnightTimestampToTime 函数
  - 重构 GetMidnightTimestamp，GetMidnightTime 实现
  - 增加北京时区常量 TIME_ZONE_BEIJING

- v1.0.6 增加参数中日期，日期时间的检查

  - ParamDateTimeCheck 日期时间参数解析
  - ParamDateCheck 日期参数解析
  - ParamDateOrDateTimeCheck 日期或日期时间参数解析
  - GormWhere 增加一组方法
    - AddDateTimeScope 日期时间范围参数解析
    - AddDateScope 日期范围参数解析
    - AddDateTimeScopeTimestamp 日期时间范围参数解析(时间戳)
    - AddDateScopeTimestamp 日期范围参数解析(时间戳)

- v1.0.5 增加一组 http 请求相关函数

  - StructToQueryParams 将带有 json 标记的结构体转换为 url.Values
  - PostRequestByOrigin 原始的 POST 请求，上传 JSON 数据并返回 JSON 响应
  - GetRequestByOrigin 原始的 Get 请求，返回 JSON 响应
  - PostRequestBy2Map 发起一个 POST 请求，上传 JSON 数据并返回 JSON 响应
  - PostRequestBy2Struct 发起一个 POST 请求，上传 JSON 数据并返回 JSON 响应
  - GetRequestByMap2Map 发起一个 GET 请求，URL 上行数据是查询参数，返回 JSON 响应
  - GetRequestByMap2Struct 发起一个 GET 请求，URL 上行数据是查询参数，返回 JSON 响应
  - GetRequestByStruct2Map 发起一个 GET 请求，URL 上行数据是查询参数，返回 JSON 响应
  - GetRequestByStruct2Struct 发起一个 GET 请求，URL 上行数据是查询参数，返回 JSON 响应

- v1.0.4
  - page 增加 Offset 和 Limit 方法(简化)
  - 增加是否是错误的类型判断
  - 增加 ParamDateTime 是否为空字符的函数
  - 增加 NewGormWhere 函数，用于简化 Gorm 的条件查询
  - 增加 GormWhere 类，用于简化 Gorm 的条件查询
  - 增加 ParamDateTime 类，用于处理日期参数
- v1.0.1

  - 增加一些单元测试
  - 增加判断集合是否相同的方法
  - 增加判断数组是否相同的方法
  - 增加数组排序的方法

- v1.0.0 正式版
  - 该版本的功能都在生产开发库中实践并提炼出来
