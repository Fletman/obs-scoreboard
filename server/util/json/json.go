package json
import (
	"errors"
	"github.com/minio/simdjson-go"
)

/* convert a byte array to a <string, object> map */
func JsonToMap(json []byte) (json_map map[string]interface{}, err error) {
	parsed_json, err := simdjson.Parse(json, nil)
	if err != nil { return }
	root_iter := parsed_json.Iter()
	if json_type := root_iter.Advance(); json_type != simdjson.TypeRoot {
		err = errors.New("Jo JSON root")
		return
	}

	json_type, json_iter, err := root_iter.Root(nil)
	if err != nil {
		return
	} else if json_type != simdjson.TypeObject {
		err = errors.New("Not a JSON object")
		return
	}

	json_object, err := json_iter.Object(nil)
	if err != nil { return }

	json_map, err = json_object.Map(nil)
	if err != nil { return }
	return
}