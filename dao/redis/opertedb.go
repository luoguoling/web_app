package redis

import (
	"context"
)
import "fmt"

func Operatedb() {
	fmt.Println("开始处理redis!!!")
	ctx := context.Background()
	err := Rdb.Set(ctx, "eeee", "ddd", 0).Err()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("数据插入成功!!!")
	}
	val, err := Rdb.Get(ctx, "eeee").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

}
