package jsondata
import "encoding/json"
import "io"
import "fmt"

type List []interface{}

func NewList() List {
	return List{}
}

func (p *List) Append(val interface{}) {
	slice := *p
	slice = append(slice, val)
	*p = slice
}

func (p * List) String() (string) {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

func (r * List) WriteTo(w io.Writer) (n int64, err error) {
	b, err := json.Marshal(r)
	if err != nil {
		return 0, err
	}
	fmt.Fprint(w, string(b))
	return int64(len(b)), nil
}

