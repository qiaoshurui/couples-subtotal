// description:
// @author renshiwei
// Date: 2022/10/13 00:43

package utils

import "time"

// SubDays 计算时间差（天数）
func SubDays(t1, t2 time.Time) (day int) {
	day = int(t1.Sub(t2).Hours() / 24)
	return
}
