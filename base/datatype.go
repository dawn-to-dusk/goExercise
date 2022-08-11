package base

import (
	"math"
)

// 常量定义
const (
	cst_a = 100
	cst_b // 100
	cst_c // 100
)

// iota
const (
	i_a      = iota               // 0
	i_b      = iota               // 1
	i_c                           // 2
	_                             // 3被忽略
	i_d                           // 4
	i_e      = 100                // 插队 5被忽略
	i_f                           // 按上一个值 6被忽略
	i_g      = iota               // 7
	i_h                           // 8
	i_i, i_j = iota + 2, iota + 3 // 9+2=11, 9+3=12
)

const (
	_  = iota             // 0
	KB = 1 << (10 * iota) // 1024
	MB = 1 << (10 * iota) // 1048576
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 支持科学计数法、支持取类型范围值
var a, b, c, d = 071, 0x1F, 1e9, math.MaxInt

var strs = `第一行
第二行
   第三行
`
