package https

import "encoding/base64"

func encodeBasicAuth(name, pwd string) string {
	return "Basic " +
		base64.URLEncoding.EncodeToString([]byte(name+":"+pwd))
}
