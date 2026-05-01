package unittest_test

import (
	unittest "go18/skills/unit_test"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 针对Add函数的单元测试
func TestAdd(t *testing.T) {
	// 只想运行一下
	// 对单元log进行配置，打印单元测试的日志
	// 如果没有打印日志，配置VsCode打印单元格的测试日志：-v -count=1
	t.Log(unittest.Add(1, 2))

	// 通过程序断言来判断
	if unittest.Add(1, 2) != 3 {
		t.Fatal("1 + 2 != 3")
	}
	
	// 专门的断言库，可以根据实际逻辑继续使用if-else嵌套
	should := assert.New(t)
	should.Equal(3, unittest.Add(1, 2))
}
