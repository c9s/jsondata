package jsondata
import "testing"


func TestJsonMap(t *testing.T) {
	resp := Map{ "success": true }
	json := resp.String()
	t.Log(json)
	if json != `{"success":true}` {
		t.Fatal(json)
	}
}



