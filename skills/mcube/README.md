# 配置（依赖注入）

## Config大对象

比较容易理解，也比较方便维护，这个配置就是和项目绑定了，但是难以被复用
在v3.0中完成了这两个映射
```go
type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
	Log         *Log         `toml:"log" yaml:"log" json:"log"`
}
```
```yaml
app:
  host: 127.0.0.1
  port: 8080
mysql:
  host: 127.0.0.1
  port: 3306
  database: go18
  username: "root"
  password: ""
  debug: true
log:
  level: debug
```

## 依赖注入

大对象时，没办法按照项目需求进行自由组装
我们想要的是分别映射，如
映射app
```yaml
app:
  host: 127.0.0.1
  port: 8080
```

映射mysql
```yaml
mysql:
  host: 127.0.0.1
  port: 3306
  database: go18
  username: "root"
  password: ""
  debug: true
```
...等等

可以导入mcube提供的方案（https://www.mcube.top/guide/component/mysql.html ）
datasource（数据源） 
```toml
[datasource]
  provider = "mysql"           # 数据库驱动，默认 mysql
  host     = "127.0.0.1"
  port     = 3306
  database = "myapp"
  username = "root"
  password = "123456"
  debug    = false             # 打印所有 SQL
  trace    = true              # 启用 OpenTelemetry 链路追踪

  # 凭证模式（可选，默认 static）
  credential_mode      = "static"
  vault_path           = ""
  vault_username_field = "username"
  vault_password_field = "password"
  vault_auto_renew     = true
  vault_renew_threshold = 0.8

  # GORM 高级选项（可选）
  skip_default_transaction = false
  prepare_stmt             = true
  dry_run                  = false
```

导入包
```go
import (

// 自动解析配置文件相应的部分
"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

func main() {
	
// 获取 *gorm.DB 对象
db := datasource.DB()
fmt.Println(db)
}
```

他会自动到配置文件去解析对应的datasource

