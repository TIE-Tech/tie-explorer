package common

import (
	"bytes"
	"math"
	"strconv"
	"time"
)

type BaseToken struct {
	TokenName     string `json:"tokenName"`
	TokenValue    string `json:"tokenValue"`
	TokenType     string `json:"tokenType"`
	TokenId       string `json:"tokenId"`
	TokenContract string `json:"tokenContract"`
}

func StrTime(atime int64) string {
	var byTime = []int64{365 * 24 * 60 * 60, 24 * 60 * 60, 60 * 60, 60, 1}
	var unit = []string{"years ago", "days ago", "hours ago", "minutes ago", "seconds ago"}
	now := time.Now().Unix()
	ct := now - atime
	if ct < 0 {
		return "just moment"
	}
	var res string
	for i := 0; i < len(byTime); i++ {
		if ct < byTime[i] {
			continue
		}
		var temp = math.Floor(float64(ct / byTime[i]))
		ct = ct % byTime[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			res = MergeString(tempStr, unit[i])
		}
		break
	}
	return res
}

func MergeString(args ...string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
	}
	return buffer.String()
}
