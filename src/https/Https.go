package https

import "net/url"

type https struct {
	uri string

	param url.Values
}
