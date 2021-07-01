package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
)

func genSign(secret string, timestamp int64) string {
	b := &[]byte{}
	*b = append(*b, strconv.FormatInt(timestamp, 10)...)
	*b = append(*b, '\n')
	*b = append(*b, secret...)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(*b)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

