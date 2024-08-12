package helper

import (
	"fmt"
	"time"
)

func InterfaceSliceToSliceInt64(payload interface{}) (res []int64) {
	dataSlice, ok := payload.([]interface{})
	if !ok {
		fmt.Println("Failed to convert to []interface{}")
		return
	}
	for _, v := range dataSlice {
		switch val := v.(type) {
		case float64:
			res = append(res, int64(val))
		default:
			return
		}
	}
	return
}

func InterfacePointerBoolToPointerBool(payload interface{}) (res *bool) {
	if payload != nil {
		if val, ok := payload.(*bool); ok {
			return val
		} else {
			return
		}
	}
	return
}

func InterfacePointerTimeToPointerTime(payload interface{}) (res *time.Time) {
	if payload != nil {
		if val, ok := payload.(*time.Time); ok {
			return val
		} else {
			return
		}
	}
	return
}
