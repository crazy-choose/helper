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
	defaultDate  = "2000-01-01"          // 默认日期
	timeParseErr = "时间解析错误: %v"
)

func CalculateAngle(pre, cur meta.KlineInfo, _fmt_ bool) float64 {
	return calcAngle(pre, cur, _fmt_)
}

func calcAngle(pre, cur meta.KlineInfo, _fmt_ bool) float64 {
	//priceChange := (k2.Close - k1.Close) / k1.Close * 100
	//disClose := cur.Close.Sub(pre.Close)
	//priceChange := disClose.Div(pre.Close).Mul(decimal.NewFromFloat(100.0))
	//fmt.Printf("上涨幅度: %v\n", priceChange)

	// 计算波动基准（方法1：两根K线平均范围）
	//a1 := ((k1.High - k1.Low) + (k2.High - k2.Low)) / 2
	avgTr := (pre.High.Sub(pre.Low)).Add(cur.High.Sub(cur.Low)).Div(decimal.NewFromInt(2))

	// 计算角度
	priceDiff := cur.Close.Sub(pre.Close)
	radians := math.Atan(priceDiff.Div(avgTr).InexactFloat64())
	angle := radians * 180 / math.Pi
	if _fmt_ {
		fmt.Printf("[%s][%s-%s] 涨|跌->角度: %v\n", cur.InstrumentId, cur.TradingDay, cur.UpdateTime, angle)
	}
	return angle
}

func calcAngle2(pre, cur meta.KlineInfo) float64 {
	// 价格变动计算
	priceDiff := cur.Close.Sub(pre.Close)
	if pre.Close.IsZero() {
		return 0
	}
	// 波动基准计算（增加除数校验）
	rangePre := pre.High.Sub(pre.Low)
	rangeCur := cur.High.Sub(cur.Low)
	avgTr := rangePre.Add(rangeCur).Div(decimal.NewFromInt(2))
	if avgTr.IsZero() {
		return 0
	}

	// 角度计算
	ratio := priceDiff.Div(avgTr)
	radians := math.Atan(ratio.InexactFloat64()) // 注意精度损失
	return radians * 180 / math.Pi
}

// target := 40.0  // 目标角度区间下限（示例值，需根据策略调整）
func AngleByClose(pre, cur meta.KlineInfo, target float64, _fmt_ bool) meta.DirectionType {
	angle := calcAngle(pre, cur, _fmt_)
	return AngleOk(angle, target)
}

// target := 40.0         // 目标角度区间下限（示例值，需根据策略调整）
// vol := 1.5             // 成交量要求：当前K线成交量 >= 前一根的1.5倍
// dur := 1 * time.Minute // 剩余时间阈值
func AngleByHalfway(pre, cur meta.KlineInfo, target, vol float64, dur time.Duration, _fmt_ bool) meta.DirectionType {
	angle := calcAngle(pre, cur, _fmt_)

	// 2. 三重过滤条件 --------------------------------------------------
	// 条件1：角度冗余度（当前角度 > 目标角度×1.2）
	angleDir := AngleOk(angle, target*1.2)
	if angleDir == meta.DirectionNone {
		if _fmt_ {
			fmt.Printf("[%s][未满足条件] 角度: %.2f度，角度冗余度: 不满足\n", cur.InstrumentId, angle)
		}
		return angleDir
	}

	// 条件2：成交量放大（当前成交量 >= 前一根×volumeRatio）
	volumeCondition := float64(cur.EVolume) >= float64(pre.EVolume)*vol
	if !volumeCondition {
		if _fmt_ {
			fmt.Printf("[%s][未满足条件] 角度: %.2f度，角度冗余度: ok，成交量放大[%v]:不满足\n", cur.InstrumentId, angle, vol)
		}
		return meta.DirectionNone
	}
	// 条件3：时间过滤（剩余时间 <= 阈值）
	timeCondition, e := calcTimeDur(cur.StartTime, cur.UpdateTime, int(cur.Bt), dur)
	if e != nil || !timeCondition {
		if _fmt_ {
			fmt.Printf("[%s][未满足条件] 角度: %.2f度，角度冗余度: ok，成交量放大[%v]:满足, 时间过滤:不满足\n", cur.InstrumentId, angle, vol)
		}
		return meta.DirectionNone
	}
	return angleDir
}

// 计算最小dur值（满足条件：start_time + bt分钟 - update_time < dur）
func calcTimeDur(startTimeStr, updateTimeStr string, bt int, dur time.Duration) (bool, error) {
	// 修正日期拼接方式
	parseWithDefaultDate := func(t string) (time.Time, error) {
		fullTimeStr := fmt.Sprintf("%s %s", defaultDate, t)
		return time.ParseInLocation(timeLayout, fullTimeStr, time.Local) // 指定本地时区
	}

	// 解析时间（自动补全默认日期）
	startTime, err := parseWithDefaultDate(startTimeStr)
	if err != nil {
		return false, fmt.Errorf(timeParseErr, err)
	}
	updateTime, err := parseWithDefaultDate(updateTimeStr)
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

func AngleOk(angle, target float64) meta.DirectionType {
	if angle > 0 {
		if angle >= target {
			return meta.DirectionBuy
		}
	} else {
		if angle <= -target {
			return meta.DirectionSell
		}
	}
	return meta.DirectionNone
}
