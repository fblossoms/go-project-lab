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
```go
func TestLoadConfigFromYaml(t *testing.T) {
	err := config.LoadConfigFromYaml(fmt.Sprintf("%sC:\\Users\\flyfl\\Desktop\\go_18\\book\\v2\\application.yaml", os.Getenv("workspaceFolder")))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("DATABASE_HOST", "DATABASE_PORT")

	err := config.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())
}
```

## 补充日志配置
```go
type Log struct {
	Level zerolog.Level `json:"level" yaml:"level" toml:"level" env:"LOG_LEVEL"`

	logger *zerolog.Logger
	lock   sync.Mutex
}

func (l *Log) SetLogger(logger zerolog.Logger) {
	l.logger = &logger
}

func (l *Log) Logger() *zerolog.Logger {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.logger == nil {
		l.SetLogger(zerolog.New(l.ConsoleWriter()).Level(l.Level).With().Caller().Timestamp().Logger())
	}

	return l.logger
}

func (c *Log) ConsoleWriter() io.Writer {
	output := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.NoColor = false
		w.TimeFormat = time.RFC3339
	})

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%-6s", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	return output
}
```