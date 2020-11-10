package model

import (
	"crypto/md5"
	"fmt"
	"github.com/snail007/gmc"
)

func OrderBy(orderBy ...interface{}) (col, by string) {
	if len(orderBy) > 0 {
		switch val := orderBy[0].(type) {
		case gmc.M:
			for k, v := range val {
				col, by = k, v.(string)
				break
			}
		case gmc.Mss:
			for k, v := range val {
				col, by = k, v
				break
			}
		}
	}
	return
}

func EncodePassword(pwd string) string {
	h:=md5.New()
	h.Write([]byte(pwd+pwd+pwd))
	return fmt.Sprintf("%x",h.Sum(nil))
}