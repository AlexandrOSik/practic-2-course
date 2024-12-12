package utils

import (
	"slices"
	"strings"
)

func GetItemIds(cookie string, cookieExists bool) []string {
	if !cookieExists {
		cookie = ""
	}
	if len(cookie) == 0 {
		return make([]string, 0)
	}
	return strings.Split(cookie, ",")
}

func GetItemNames(cookie string, cookieExists bool) []string {
	ids := GetItemIds(cookie, cookieExists)
	for i := 0; i < len(ids); i++ {
		ids[i] = Const.KEY_PAGE_ITEMS[slices.Index(Const.KEY_PAGE_SYMBOL, ids[i])]
	}
	return ids
}

func GetItemExistance(cookie string, cookieExists bool) map[string]bool {
	result := make(map[string]bool)
	ids := GetItemIds(cookie, cookieExists)

	for i := 0; i < len(Const.KEY_PAGE_ITEMS); i++ {
		result[Const.KEY_PAGE_ITEMS[i]] = slices.Contains(ids, Const.KEY_PAGE_SYMBOL[i])
	}

	return result
}
