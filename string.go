package commons_go_lang

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	EMPTY_STRING = ""
)

//
//
// StringIgnoreEquals is ignore to compare
//
//
func StringIgnoreEquals(src string, dest string) bool {
	src = strings.ToLower(src)
	dest = strings.ToLower(dest)
	return src == dest
}

//
//
// StringSplitByRegexp is use regexp to split string
//
//
func StringSplitByRegexp(src string, reg string) []string {
	pattern := regexp.MustCompile(reg)
	indexes := pattern.FindAllStringIndex(src, -1)
	if len(indexes) == 0 {
		return nil
	}
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = src[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = src[laststart:]
	return result
}

//
//
// StringFirstLetterLower
//
//
func StringFirstLetterLower(src string) string {
	return strings.ToLower(src[:1]) + src[1:]
}

//
//
// StringFirstLetterUpper
//
//
func StringFirstLetterUpper(src string) string {
	return strings.ToUpper(src[:1]) + src[1:]
}

//
//
// StringBetween is get string in between start and end
//
//
func StringBetween(src string, start string, end string) (string, error) {
	sI := strings.Index(src, start)
	if sI >= 0 {
		src = src[sI+len(start):]
		eI := strings.Index(src, end)
		if eI > 0 {
			return src[:eI], nil
		} else {
			return src, errors.New("not found the end " + end)
		}
	} else {
		return EMPTY_STRING, errors.New("not fround the start " + start)
	}
}

//
//
// StringStartWith is check if start with
//
//
func StringStartWith(src string, s string) bool {
	return strings.Index(src, s) == 0
}

//
//
// StringAfter get string after give string
//
//
func StringAfter(src string, start string) (string, bool) {
	sI := strings.Index(src, start)
	if sI >= 0 {
		return src[sI+len(start):], true
	} else {
		return EMPTY_STRING, false
	}
}

//
//
// StringBefore get string before give string
//
//
func StringBefore(src string, start string) (string, bool) {
	sI := strings.Index(src, start)
	if sI >= 1 {
		return src[:sI], true
	} else {
		return EMPTY_STRING, false
	}
}

//
//
// StringMatch is get regexp value by group index
//
//
func StringMatch(src string, p string, group int) (string, error) {
	pattern, err := regexp.Compile(p)
	if err != nil {
		return EMPTY_STRING, err
	} else {
		r := pattern.FindStringSubmatch(src)
		if len(r) > group {
			return r[group], nil
		} else {
			return EMPTY_STRING, errors.New(fmt.Sprintf("not found group %d", group))
		}
	}
}

//
//
// StringIsEmpty is check if empty
//
//
func StringIsEmpty(src string) bool {
	return strings.Trim(src, " ") == ""
}

//
//
// StringWrap
//
//
func StringWrap(src string, wrap string) string {
	return wrap + src + wrap
}
