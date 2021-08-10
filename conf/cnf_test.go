package conf

import (
	"fmt"
	"testing"
)

func TestCal(t *testing.T) {
	// 等级最低奖金 + (实际销售额-当前等级门槛) / (下一级门槛-当前等级门槛) * (下一级最低奖金-当前等级最低奖金)
	result1 := 2000 + (19.5-19)/(40-19) * (4000-2000)
	result2 := 2000 + (39.5-19)/(40-19) * (4000-2000)

	fmt.Println(result1)
	fmt.Println(result2)
}