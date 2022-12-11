package biz

import (
	"strings"
	"time"
)

type PointCond struct {
	Begin int
	Count int
	PIDs  []int32
	UIDs  []int32
}

func (c *PointCond) ParseCond() (whereStage string, args []interface{}) {
	var builder strings.Builder
	if len(c.PIDs) != 0 {
		builder.WriteString("`pid` IN (?)")
		args = append(args, c.PIDs)
	}
	if len(c.UIDs) != 0 {
		if builder.Len() != 0 {
			builder.WriteString(" AND ")
		}
		builder.WriteString("`uid` IN (?)")
		args = append(args, c.UIDs)
	}
	return builder.String(), args
}

type RecordCond struct {
	Begin        int
	Count        int
	RIDs         []int32
	PIDs         []int32
	MinClickedAt int64
	MaxClickedAt int64
}

func (c *RecordCond) ParseCond() (whereStage string, args []interface{}) {
	var builder strings.Builder
	if len(c.RIDs) != 0 {
		builder.WriteString("`rid` IN (?)")
		args = append(args, c.RIDs)
	}
	if len(c.PIDs) != 0 {
		if builder.Len() != 0 {
			builder.WriteString(" AND ")
		}
		builder.WriteString("`pid` IN (?)")
		args = append(args, c.PIDs)
	}
	if c.MinClickedAt != 0 {
		if builder.Len() != 0 {
			builder.WriteString(" AND ")
		}
		builder.WriteString("`clicked_at` > ?")
		args = append(args, time.Unix(c.MinClickedAt, 0))
	}
	if c.MaxClickedAt != 0 {
		if builder.Len() != 0 {
			builder.WriteString(" AND ")
		}
		builder.WriteString("`clicked_at` < ?")
		args = append(args, time.Unix(c.MaxClickedAt, 0))
	}
	return builder.String(), args
}
