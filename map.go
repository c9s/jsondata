package jsondata
import "encoding/json"
import "bytes"
import "fmt"
import "io"

type Map map[string]interface{}

func New() Map {
	return Map{}
}

func NewMap() Map {
	return Map{}
}

func (r * Map) Debug() string {
	str := r.String()
	newBytes := new(bytes.Buffer)
	err := json.Indent(newBytes, []byte(str), "", "  ")
	if err != nil {
		return ""
	}
	return newBytes.String()
}

func (r * Map) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}

func (r * Map) WriteTo(w io.Writer) (n int64, err error) {
	b, err := json.Marshal(r)
	if err != nil {
		return 0, err
	}
	fmt.Fprint(w, string(b))
	return int64(len(b)), nil
}

