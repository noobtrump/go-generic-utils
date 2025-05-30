// Copyright 2024 Humphrey
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tuple

import (
	"strconv"
	"testing"
)

// 准备基准测试数据
func prepareIntData(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return data
}

func prepareStringData(n int) []string {
	data := make([]string, n)
	for i := 0; i < n; i++ {
		data[i] = "str" + strconv.Itoa(i)
	}
	return data
}

// 基准测试: 创建Pair对象
func BenchmarkNewPair(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for j := 0; j < size; j++ {
					_ = NewPair(keys[j], values[j])
				}
			}
		})
	}
}

// 基准测试: 从键值数组创建Pairs
func BenchmarkNewPairs(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = NewPairs(keys, values)
			}
		})
	}
}

// 基准测试: 将Pairs拆分为键值数组
func BenchmarkSplitPairs(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)
			pairs, _ := NewPairs(keys, values)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = SplitPairs(pairs)
			}
		})
	}
}

// 基准测试: 将Pairs展平为扁平数组
func BenchmarkFlattenPairs(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)
			pairs, _ := NewPairs(keys, values)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = FlattenPairs(pairs)
			}
		})
	}
}

// 基准测试: 将扁平数组打包为Pairs
func BenchmarkPackPairs(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)
			pairs, _ := NewPairs(keys, values)
			flatPairs := FlattenPairs(pairs)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// 注意: 这里可能会有panic，因为PackPairs需要进行类型断言
				// 在基准测试中，我们会尝试恢复panic以避免测试中断
				func() {
					defer func() {
						if r := recover(); r != nil {
							// 恢复panic
						}
					}()
					_ = PackPairs[int, string](flatPairs)
				}()
			}
		})
	}
}

// 基准测试: 创建Triple对象
func BenchmarkNewTriple(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			ints := prepareIntData(size)
			strings := prepareStringData(size)
			floats := make([]float64, size)
			for i := 0; i < size; i++ {
				floats[i] = float64(i) + 0.5
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for j := 0; j < size; j++ {
					_ = NewTriple(ints[j], strings[j], floats[j])
				}
			}
		})
	}
}

// 基准测试: 从Map创建Pairs
func BenchmarkPairsFromMap(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			m := make(map[int]string, size)
			for i := 0; i < size; i++ {
				m[i] = "str" + strconv.Itoa(i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = PairsFromMap(m)
			}
		})
	}
}

// 基准测试: 从Pairs创建Map
func BenchmarkMapFromPairs(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)
			pairs, _ := NewPairs(keys, values)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = MapFromPairs(pairs)
			}
		})
	}
}

// 基准测试: Range遍历Pairs
func BenchmarkRange(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)
			pairs, _ := NewPairs(keys, values)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Range(pairs, func(k int, v string) error {
					return nil
				})
			}
		})
	}
}

// 基准测试: Filter过滤Pairs
func BenchmarkFilter(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)
			pairs, _ := NewPairs(keys, values)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// 过滤偶数键
				_ = Filter(pairs, func(k int, v string) bool {
					return k%2 == 0
				})
			}
		})
	}
}

// 基准测试: Map转换Pairs
func BenchmarkMapTransform(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)
			pairs, _ := NewPairs(keys, values)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// 将键翻倍，值转为大写
				_ = Map(pairs, func(k int, v string) (int, string) {
					return k * 2, v + "_mapped"
				})
			}
		})
	}
}

// 基准测试: Reduce归约Pairs
func BenchmarkReduce(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			keys := prepareIntData(size)
			values := prepareStringData(size)
			pairs, _ := NewPairs(keys, values)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// 计算所有键的总和
				_ = Reduce(pairs, 0, func(r int, k int, v string) int {
					return r + k
				})
			}
		})
	}
}

// 基准测试: 商品价格元组操作 (电商场景)
func BenchmarkProductPriceOperations(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			// 创建商品价格列表
			prices := make(ProductPriceList, size)
			for i := 0; i < size; i++ {
				prices[i] = NewProductPrice("P"+strconv.Itoa(i), float64(i*10)+0.99)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// 排序
				pricesCopy := make(ProductPriceList, len(prices))
				copy(pricesCopy, prices)
				pricesCopy.SortByPrice()

				// 过滤价格区间
				filtered := pricesCopy.FilterByPriceRange(50, 200)

				// 计算总价
				_ = filtered.TotalPrice()
			}
		})
	}
}

// 基准测试: 购物车项目操作
func BenchmarkCartItemOperations(b *testing.B) {
	sizes := []int{5, 10, 20, 50}

	for _, size := range sizes {
		b.Run("size="+strconv.Itoa(size), func(b *testing.B) {
			// 创建购物车项目列表
			cart := make(CartItemList, size)
			for i := 0; i < size; i++ {
				quantity := (i % 5) + 1 // 1-5件
				price := float64(i*10) + 9.99
				cart[i] = NewCartItem("P"+strconv.Itoa(i), quantity, price)
				// 随机设置一些项目为未选中
				if i%3 == 0 {
					cart[i].Selected = false
				}
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// 计算选中项目总量
				_ = cart.TotalQuantity()

				// 计算选中项目总金额
				_ = cart.TotalAmount()

				// 过滤选中项目
				_ = cart.FilterSelected()

				// 更新某项目数量
				cart.UpdateQuantity("P3", 10)
			}
		})
	}
}
