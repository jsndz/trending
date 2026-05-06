package util

import (
	"net/url"
	"strings"
)

func NormalizeURL(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	u.Scheme = strings.ToLower(u.Scheme)
	u.Host = strings.ToLower(u.Host)

	u.Host = strings.TrimSuffix(u.Host, ":80")
	u.Host = strings.TrimSuffix(u.Host, ":443")

	u.Path = strings.TrimRight(u.Path, "/")

	q := u.Query()

	for key := range q {
		k := strings.ToLower(key)

		if strings.HasPrefix(k, "utm_") ||
			k == "fbclid" ||
			k == "gclid" {
			q.Del(key)
		}
	}

	u.RawQuery = q.Encode()

	u.Fragment = ""

	return u.String(), nil
}
