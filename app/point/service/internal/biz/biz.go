package biz

import (
	"github.com/google/wire"
	"strings"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewPointUseCase, NewRecordUseCase, NewTotalUseCase)

type PointCond struct {
	Begin uint32
	Count uint32
	PIDs  []uint32
	UIDs  []uint32
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
	Begin        uint32
	Count        uint32
	RIDs         []uint32
	PIDs         []uint32
	MinClickedAt uint64
	MaxClickedAt uint64
}

func (c *RecordCond) ParseCond() (whereStage string, args []interface{}) {
	var whereStages []string
	if len(c.RIDs) != 0 {
		whereStages = append(whereStages, "`rid` IN (?)")
		args = append(args, c.RIDs)
	}
	if len(c.PIDs) != 0 {
		whereStages = append(whereStages, "`pid` IN (?)")
		args = append(args, c.PIDs)
	}
	if c.MinClickedAt != 0 {
		whereStages = append(whereStages, "`clicked_at` >= ?")
		args = append(args, time.Unix(int64(c.MinClickedAt), 0))
	}
	if c.MaxClickedAt != 0 {
		whereStages = append(whereStages, "`clicked_at` <= ?")
		args = append(args, time.Unix(int64(c.MaxClickedAt), 0))
	}
	return strings.Join(whereStages, " AND "), args
}
