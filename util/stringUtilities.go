package util

import "strings"

func RemoveFirstUrlPathElement(url string) string {
	//TODO do regex?
	firstIndex := strings.Index(url, "/")
	url = url[firstIndex+1:]
	secondIndex := strings.Index(url, "/")
	url = url[secondIndex:]
	return url
}
