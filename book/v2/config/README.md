# 程序的配置管理
## 程序内部可访问
```go
// 做成全局变量
var config *Config
```

## 配置的加载
```go
// 用于加载配置
config.LoadConfigFromYaml(yamlConfigFilePath)
```

## 程序配置如何使用加载进来的配置
```go
// C为Get Config缩写
// 返回ConfigObject
config.C().MySQL.Host 
```

## 为包添加单元测试
如何验证我们这个包的业务逻辑是否正确