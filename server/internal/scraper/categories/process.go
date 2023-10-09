package categories

import "net/url"

func processHref(s string) string {
	fullUrl, err := url.Parse(s)

	if err != nil {
		return s + "?action=render"
	} else {
		return BASE_URL + fullUrl.Path + "?action=render"
	}
}
