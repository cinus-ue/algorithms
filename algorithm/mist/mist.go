/**
MIT License

Copyright (c) 2020 AsyncIns

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package mist

/*
* 薄雾算法
*
* 1      2                                                     48         56       64
* +------+-----------------------------------------------------+----------+----------+
* retain | increase                                             | salt     | sequence |
* +------+-----------------------------------------------------+----------+----------+
* 0      | 0000000000 0000000000 0000000000 0000000000 0000000 | 00000000 | 00000000 |
* +------+-----------------------------------------------------+------------+--------+
*
* 0. 最高位，占 1 位，保持为 0，使得值永远为正数；
* 1. 自增数，占 47 位，自增数在高位能保证结果值呈递增态势，遂低位可以为所欲为；
* 2. 随机因子一，占 8 位，上限数值 255，使结果值不可预测；
* 3. 随机因子二，占 8 位，上限数值 255，使结果值不可预测；
*
* 编号上限为百万亿级，上限值计算为 140737488355327 即 int64(1 << 47 - 1)，假设每天取值 10 亿，能使用 385+ 年
 */

import (
	"crypto/rand"
	"math/big"
	"sync"
)

const saltBit = uint(8)                   // 随机因子二进制位数
const saltShift = uint(8)                 // 随机因子移位数
const increaseShift = saltBit + saltShift // 自增数移位数

type Mist struct {
	sync.Mutex       // 互斥锁
	increase   int64 // 自增数
	saltA      int64 // 随机因子一
	saltB      int64 // 随机因子二
}

/* 初始化 Mist 结构体*/
func NewMist() *Mist {
	mist := Mist{increase: 1}
	return &mist
}

/* 生成唯一编号 */
func (c *Mist) Generate() int64 {
	c.Lock()
	c.increase++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	randA, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltA = randA.Int64()
	randB, _ := rand.Int(rand.Reader, big.NewInt(255))
	c.saltB = randB.Int64()
	// 通过位运算实现自动占位
	mist := (c.increase << increaseShift) | (c.saltA << saltShift) | c.saltB
	c.Unlock()
	return mist
}
