package payments

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/url"
	"sort"
	"strings"
)

// getConcatenationParams concatenation of pairs parameter name=parameter
// value, placed in ascending alphabetical order according to the
// parameter name.
func getConcatenationParams(values url.Values) string {
	if values == nil {
		return ""
	}

	var buf strings.Builder

	keys := make([]string, 0, len(values))

	for k := range values {
		if k == "sig" {
			continue
		}

		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		vs := values[k]
		for _, v := range vs {
			_, _ = buf.WriteString(k)
			_ = buf.WriteByte('=')
			_, _ = buf.WriteString(v)
		}
	}

	return buf.String()
}

// toHex return hex string.
func toHex(src []byte) string {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	return string(dst)
}

// Sign return signature.
//
// The parameter sig equals md5 from the concatenation of pairs parameter
// name=parameter value, placed in ascending alphabetical order according
// to the parameter name and the secret signature api_secret indicated in
// your app settings.
func Sign(values url.Values, secret string) string {
	h := md5.New()
	_, _ = io.WriteString(h, getConcatenationParams(values)+secret)

	return toHex(h.Sum(nil))
}
