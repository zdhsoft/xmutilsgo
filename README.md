# xmutilsgo

## 使用办法

```bash
go get github.com/zdhsoft/xmutilsgo
```

## 主要内容

- int.go 定义泛型的整数类型和字符串转整数的函数和随机范围的函数
- isin.go 判断指定元素是否再数组中的函数
- page.go mysql用于分页的类
- ret.go 通用返回值的类
- set.go 基于map实现的集合功能
- string.go 主要是字符串链接和pad的函数
- time.go 用于时间处理的功能函数
- array.go 数组相关工具函数
- map.go map相关工具函数
- json.go json相关工具函数
- md5.go md5相关工具函数
- utils.go 其他工具函数
- version.go 版本信息
- request.go 封装了一些常用的http请求函数

## 版本信息

- v1.0.7 增加一些0点计算的函数
  - 增加GetMidnightTimeToTime和GetMidnightTimestampToTime函数
  - 重构GetMidnightTimestamp，GetMidnightTime实现
  - 增加北京时区常量TIME_ZONE_BEIJING

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
  - StructToQueryParams 将带有json标记的结构体转换为url.Values
  - PostRequestByOrigin 原始的POST 请求，上传 JSON 数据并返回 JSON 响应
  - GetRequestByOrigin 原始的Get 请求，返回 JSON 响应
  - PostRequestBy2Map 发起一个 POST 请求，上传 JSON 数据并返回 JSON 响应
  - PostRequestBy2Struct 发起一个 POST 请求，上传 JSON 数据并返回 JSON 响应
  - GetRequestByMap2Map 发起一个 GET 请求，URL 上行数据是查询参数，返回 JSON 响应
  - GetRequestByMap2Struct 发起一个 GET 请求，URL 上行数据是查询参数，返回 JSON 响应
  - GetRequestByStruct2Map 发起一个 GET 请求，URL 上行数据是查询参数，返回 JSON 响应
  - GetRequestByStruct2Struct 发起一个 GET 请求，URL 上行数据是查询参数，返回 JSON 响应

- v1.0.4
  - page增加Offset和Limit方法(简化)
  - 增加是否是错误的类型判断
  - 增加ParamDateTime是否为空字符的函数
  - 增加NewGormWhere函数，用于简化Gorm的条件查询
  - 增加GormWhere类，用于简化Gorm的条件查询
  - 增加ParamDateTime类，用于处理日期参数
- v1.0.1
  - 增加一些单元测试
  - 增加判断集合是否相同的方法
  - 增加判断数组是否相同的方法
  - 增加数组排序的方法

- v1.0.0 正式版
  - 该版本的功能都在生产开发库中实践并提炼出来
