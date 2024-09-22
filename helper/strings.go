package helper

import "strings"

func SafeStringQ(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Replace(str, `<`, `&lt;`, -1)
	str = strings.Replace(str, `>`, `&gt;`, -1)
	str = strings.Replace(str, `'`, `&apos;`, -1)
	str = strings.Replace(str, `"`, `&quot;`, -1)
	str = strings.Replace(str, `\`, `\\`, -1)
	return `'` + str + `'`
}