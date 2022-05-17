package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, `

Respected user:
We will not collect any of your information, and any right to dispose of the data belongs to you.

尊敬的用户:
我们不会搜集您的任何信息,对数据的任何处置权都属于您自己.

尊敬的用戶:
我們不會蒐集您的任何信息,對數據的任何處置權都屬於您自己.

2022.3.1 Reading
	`)
}
