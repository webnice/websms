package response

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"bytes"
)

// Отображение ответа в дебаг режиме
func (r *Response) debugResponse(data []byte) {
	const prefixKey = `< `
	var buf []byte
	var tmp [][]byte
	var i int

	defer func() { _ = recover() }()
	tmp, buf = bytes.Split(data, []byte("\n")), buf[:0]
	for i = range tmp {
		tmp[i] = bytes.TrimRight(tmp[i], "\r")
		buf = bytes.Join([][]byte{buf, []byte(prefixKey), tmp[i], []byte("\r\n")}, []byte(``))
	}
	r.debugFunc(bytes.TrimRight(buf, "\r\n"))
}
