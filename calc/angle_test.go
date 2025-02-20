package calc

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// 定义K线结构体（新增时间、成交量字段）
type KLine struct {
	High        float64
	Low         float64
	Close       float64
	Time        time.Time     // K线开始时间
	ElapsedTime time.Duration // 当前K线已过去的时间
	TotalTime   time.Duration // K线总周期（如5分钟）
	Volume      int64         // 成交量
}

func TestAngle(t *testing.T) {
	// 定义两根K线数据
	k1 := KLine{High: 2885, Low: 2877, Close: 2885}
	k2 := KLine{High: 2911, Low: 2883, Close: 2903}

	// 计算上涨幅度
	//priceChange := (k2.Close - k1.Close) / k1.Close * 100
	//fmt.Printf("上涨幅度: %.2f%%\n", priceChange)

	// 计算波动基准（方法1：两根K线平均范围）
	a1 := ((k1.High - k1.Low) + (k2.High - k2.Low)) / 2

	// 计算角度
	priceDiff := k2.Close - k1.Close
	radians := math.Atan(priceDiff / a1)
	angle := radians * 180 / math.Pi
	fmt.Printf("上涨角度: %.2f度\n", angle)
}

func TestAngle2(t *testing.T) {
	// 定义两根K线数据

	k1 := KLine{High: 2897, Low: 2891, Close: 2897}
	k2 := KLine{High: 2897, Low: 2882, Close: 2888}
	// 计算上涨幅度
	priceChange := (k2.Close - k1.Close) / k1.Close * 100
	fmt.Printf("上涨幅度: %.2f%%\n", priceChange)

	// 计算波动基准（方法1：两根K线平均范围）
	a1 := ((k1.High - k1.Low) + (k2.High - k2.Low)) / 2

	// 计算角度
	priceDiff := k2.Close - k1.Close
	radians := math.Atan(priceDiff / a1)
	angle := radians * 180 / math.Pi
	fmt.Printf("上涨角度: %.2f度\n", angle)
}

func TestAngle3(t *testing.T) {
	// 定义两根K线数据（带时间、成交量）
	k1 := KLine{
		High:        2885,
		Low:         2877,
		Close:       2885,
		Time:        time.Date(2023, 1, 1, 9, 30, 0, 0, time.UTC), // 假设K线开始时间
		ElapsedTime: 5 * time.Minute,                              // 第一根K线已结束
		TotalTime:   5 * time.Minute,                              // 5分钟K线
		Volume:      10000,                                        // 成交量1万手
	}
	k2 := KLine{
		High:        2911,
		Low:         2883,
		Close:       2903,
		Time:        time.Date(2023, 1, 1, 9, 35, 0, 0, time.UTC), // 第二根K线开始时间
		ElapsedTime: 4 * time.Minute,                              // 已运行4分钟（剩余1分钟）
		TotalTime:   5 * time.Minute,
		Volume:      18000, // 成交量1.8万手（比前一根放大80%）
	}

	// 1. 计算基础指标 --------------------------------------------------
	priceChange := (k2.Close - k1.Close) / k1.Close * 100
	a1 := ((k1.High - k1.Low) + (k2.High - k2.Low)) / 2
	priceDiff := k2.Close - k1.Close
	radians := math.Atan(priceDiff / a1)
	angle := radians * 180 / math.Pi

	// 2. 三重过滤条件 --------------------------------------------------
	targetAngle := 40.0              // 目标角度区间下限（示例值，需根据策略调整）
	volumeRatio := 1.5               // 成交量要求：当前K线成交量 >= 前一根的1.5倍
	timeThreshold := 1 * time.Minute // 剩余时间阈值

	// 条件1：角度冗余度（当前角度 > 目标角度×1.2）
	angleCondition := angle > targetAngle*1.2

	// 条件2：成交量放大（当前成交量 >= 前一根×volumeRatio）
	volumeCondition := float64(k2.Volume) >= float64(k1.Volume)*volumeRatio

	// 条件3：时间过滤（剩余时间 <= 阈值）
	remainingTime := k2.TotalTime - k2.ElapsedTime
	timeCondition := remainingTime <= timeThreshold

	// 3. 综合判断是否入场 ----------------------------------------------
	if angleCondition && volumeCondition && timeCondition {
		fmt.Printf("[满足条件] 入场信号触发！当前角度: %.2f度，成交量放大: %d→%d，剩余时间: %v\n",
			angle, k1.Volume, k2.Volume, remainingTime)
	} else {
		fmt.Printf("[未满足条件] 角度: %.2f度，成交量: %d→%d，剩余时间: %v\n",
			angle, k1.Volume, k2.Volume, remainingTime)
	}

	// 输出基础指标
	fmt.Printf("上涨幅度: %.2f%%\n", priceChange)
	fmt.Printf("上涨角度: %.2f度\n", angle)
}
