package calc

import (
	"fmt"
	"github.com/crazy-choose/helper/model/meta"
	"github.com/shopspring/decimal"
	"math"
	"time"
)

// 定义时间格式和固定日期（用于解析不含日期的时间字符串）
const (
	timeLayout   = "2006-01-02 15:04:05" // Go标准时间格式
	defaultDate  = "2000-01-01 "         // 默认日期
	timeParseErr = "时间解析错误: %v"
)

func AngleByClose(pre, cur meta.KlineInfo) float64 {
	//priceChange := (k2.Close - k1.Close) / k1.Close * 100
	disClose := cur.Close.Sub(pre.Close)
	priceChange := disClose.Div(pre.Close).Mul(decimal.NewFromFloat(100.0))
	fmt.Printf("上涨幅度: %v\n", priceChange)

	// 计算波动基准（方法1：两根K线平均范围）
	//a1 := ((k1.High - k1.Low) + (k2.High - k2.Low)) / 2
	avgTr := (pre.High.Sub(pre.Low)).Add(cur.High.Sub(cur.Low)).Div(decimal.NewFromInt(2))

	// 计算角度
	priceDiff := cur.Close.Sub(pre.Close)
	radians := math.Atan(priceDiff.Div(avgTr).InexactFloat64())
	angle := radians * 180 / math.Pi
	return angle
}

// target := 40.0         // 目标角度区间下限（示例值，需根据策略调整）
// vol := 1.5             // 成交量要求：当前K线成交量 >= 前一根的1.5倍
// dur := 1 * time.Minute // 剩余时间阈值
func AngleByHalfway(pre, cur meta.KlineInfo, target, vol float64, dur time.Duration) bool {
	angle := AngleByClose(pre, cur)

	// 2. 三重过滤条件 --------------------------------------------------
	// 条件1：角度冗余度（当前角度 > 目标角度×1.2）
	angleCondition := angle > target*1.2

	// 条件2：成交量放大（当前成交量 >= 前一根×volumeRatio）
	volumeCondition := float64(cur.EVolume) >= float64(pre.EVolume)*vol

	// 条件3：时间过滤（剩余时间 <= 阈值）
	timeCondition, e := calcTimeDur(cur.StartTime, cur.UpdateTime, int(cur.Bt), dur)
	if e != nil || !timeCondition {
		fmt.Printf("[%s][未满足条件] 角度: %.2f度，角度冗余度:%v，成交量放大:%v, 剩余时间: %v\n", cur.InstrumentId, angle, angleCondition, volumeCondition, e)
		return false
	}
	fmt.Printf("[%s][满足条件] 角度: %.2f度，角度冗余度:%v，成交量放大:%v, 剩余时间:%v \n", cur.InstrumentId, angle, angleCondition, volumeCondition, timeCondition)
	return angleCondition && volumeCondition && timeCondition
}

// 计算最小dur值（满足条件：start_time + bt分钟 - update_time < dur）
func calcTimeDur(startTimeStr, updateTimeStr string, bt int, dur time.Duration) (bool, error) {
	// 解析时间（自动补全默认日期）
	startTime, err := time.Parse(timeLayout, defaultDate+startTimeStr)
	if err != nil {
		return false, fmt.Errorf(timeParseErr, err)
	}
	updateTime, err := time.Parse(timeLayout, defaultDate+updateTimeStr)
	if err != nil {
		return false, fmt.Errorf(timeParseErr, err)
	}

	// 计算时间变化
	endTime := startTime.Add(time.Duration(bt) * time.Minute)
	diff := endTime.Sub(updateTime) // 时间差（可能是正/负）

	// 最小dur逻辑：
	// - 若 newTime 在 updateTime 之后（diff > 0），则 dur必须 > diff
	// - 若 newTime 在 updateTime 之前（diff <= 0），则任意 dur > diff 即可（通常取0）
	return diff < dur, nil // 若diff为负，
}
