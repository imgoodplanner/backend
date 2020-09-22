package patch

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Map represents a mapping between structure tag values and their corresponding struct field names
// Input map is set using "From" methods, Output map is set using "To" methods
// patch.Map are generated by a global mapper, which is configured to look for specific struct tag for input and output
// As an example, a patch.Map of this model (with a mapper using json & sql tags):
//
// type TestModel struct {
//      First string `json:"json_first,omitempty" sql:"db_first"`
//      Second string `json:"json_second" sql:"db_second"`
//      Third string
//      Fourth string `sql:"db_fourth"`
//      Fifth string `json:"json_fifth"`
// }
//
// would have an input {"json_first": "First", "json_second": "Second", "fifth_field": "Fifth"}
// and a output {"First": "db_first", "Second": "db_second", "Fourth": `db_fourth`}
// The interest of this struct is to handle resources patching from a request to final storer
type Map struct {
	input   map[string]string
	output  map[string]string
	Invalid bool
}

// FromHTTP decode a HTTP request into a ready-to-be-used patch.Map
// It does define JSON keys found inside request body as input map
// It ignores errors if faced, considering it should be handled outside of this function
// It instantiates later used output map and it does restaure body's io.Reader after consuming it
func FromHTTP(req *http.Request) Map {
	// prepare body reader to be cloned
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(req.Body)

	// decode input body into map of string/string
	jsonInput := make(map[string]string)
	// if an err occurs, we'll return a nil json map
	_ = json.Unmarshal(buf.Bytes(), &jsonInput)

	// reset body after we used it
	req.Body = ioutil.NopCloser(buf)

	emptyJSON := false
	if len(jsonInput) == 0 {
		emptyJSON = true
	}

	return Map{
		input:   jsonInput,
		output:  make(map[string]string),
		Invalid: emptyJSON,
	}
}

// ToModel does an intersection between current input map and given model input map
// The model patch.Map is get from mapper
// It builds current output map using model output map, matching input intersection
func (p Map) ToModel(model interface{}) Map {
	// get map for asked model
	modelMap := mapper.loadOrInitMap(model)

	// detect if we receive unsupported patch
	for k := range p.input {
		_, ok := modelMap.input[k]
		if !ok {
			p.Invalid = true
			return p
		}
	}

	// perform the intersect between model map input and current map input
	// 1. we set values from model map to the current map if there is a match
	for k := range modelMap.input {
		_, ok := p.input[k]
		if ok {
			fieldName := modelMap.input[k]
			p.input[k] = fieldName
			// 2. we also add output if it does exist
			outputTag, ok := modelMap.output[fieldName]
			if ok {
				p.output[fieldName] = outputTag
			}
		}
	}
	return p
}

// Info converts a patch.Map into a patch.Info
func (p Map) GetInfo(model interface{}) Info {
	inputKeys, outputKeys := p.getKeys()
	return Info{
		Input:  inputKeys,
		Output: outputKeys,
		Model:  model,
	}
}

// getKeys return input and output keys
func (p Map) getKeys() ([]string, []string) {
	inputKeys := []string{}
	outputKeys := []string{}
	for iPos, iVal := range p.input {
		for _, oVal := range p.output {
			if iPos == oVal {
				inputKeys = append(inputKeys, iVal)
				outputKeys = append(outputKeys, oVal)
				break
			}
		}
	}
	return inputKeys, outputKeys
}
