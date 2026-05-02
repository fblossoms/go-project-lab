package config

import (
	"fmt"
	"go18/book/v3/models"
	"sync"

	"github.com/infraboard/mcube/v2/tools/pretty"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Default() *Config {
	return &Config{
		Application: &application{
			Host: "127.0.0.1",
			Port: 8080,
		},
		MySQL: &mySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "go18",
			Username: "root",
			Password: "",
			Debug:    true,
		},
	}
}

// Config
// 定义整个程序的配置对象
// 通过config.Config来获取
// 不同标签对应不同的工具，例如是json就使用json工具
type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
}

func (c *Config) String() string {
	return pretty.ToJSON(c)
}

type application struct {
	Host string `toml:"host" yaml:"host" json:"host" env:"HOST"`
	Port int    `toml:"port" yaml:"port" json:"port" env:"PORT"`
}

type mySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"unit_test" yaml:"unit_test" toml:"unit_test" env:"DATASOURCE_DEBUG"`

	// 维持一个连接池：gorm连接对象，只允许有一个，不允许重复生成（减轻对数据库的压力），最好的方法是加互斥锁
	db *gorm.DB
	// 加互斥锁
	lock sync.Mutex
}

// GetDB 抽象出对象，全局共享
// 基本原则：所有对对象的数据主要来源于哪个数据结构，就把他变成哪个数据结构的方法
func (m *mySQL) GetDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.DB == "" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s?parseTime=true&loc=Local",
			m.Username,
			m.Password,
			m.Host,
			m.Port,
			m.DB,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db.AutoMigrate(&models.Book{})
		m.db = db
	}

	return m.db
}
