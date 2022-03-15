package log

import (
	"fmt"
	"time"
)

type TextColorFormatter struct {
	IgnoreBasicFields bool
}

func (f *TextColorFormatter) Format(e *Entry) error {
	if !f.IgnoreBasicFields {
		// e.Buffer.WriteString(fmt.Sprintf("%s %s", e.Time.Format(time.RFC3339), LevelNameMapping[e.Level]))
		e.Buffer.WriteString(fmt.Sprintf("%c[1;40;32m%s%c[0m", 0x1B, fmt.Sprintf("%s %s", e.Time.Format(time.RFC3339), LevelNameMapping[e.Level]), 0x1B))
		if e.File != "" {
			short := e.File
			for i := len(e.File) - 1; i > 0; i-- {
				if e.File[i] == '/' {
					short = e.File[i+1:]
					break
				}
			}
			// 输出内容 2022-01-04T12:21:03+08:00 INFO example.go:11
			// e.Buffer.WriteString(fmt.Sprintf(" %s:%d",short, e.Line))
			e.Buffer.WriteString(fmt.Sprintf("%c[1;40;32m%s%c[0m", 0x1B, fmt.Sprintf(" %s:%d",short, e.Line), 0x1B))
		}
		e.Buffer.WriteString(" ")
	}

	switch e.Format {
	case FmtEmptySeparate:
		// e.Buffer.WriteString(fmt.Sprint(e.Args))
		// 格式：%c[显示方式;前景色;背景色m
		// 其中0x1B是标记，%c是颜色标记，相当于固定格式[开始定义颜色，1代表高亮，40代表黑色背景，32代表绿色前景，0代表恢复默认颜色
		e.Buffer.WriteString(fmt.Sprintf("%c[1;40;32m%s%c[0m",0x1B, e.Args, 0x1B))
	default:
		// e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args))
		e.Buffer.WriteString(fmt.Sprintf("%c[1;40;32m%s%c[0m" + e.Format, 0x1B, e.Args, 0x1B))
	}
	e.Buffer.WriteString("\n")

	return nil
}
