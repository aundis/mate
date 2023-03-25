package meta

type Slot struct{}

type Listen struct{}

type ObjectMeta struct {
	Name      string          `json:"name"`
	Kind      string          `json:"kind"`
	Functions []*FunctionMeta `json:"functions"`
}

type FunctionMeta struct {
	Name       string
	Parameters []*FieldMeta
	Results    []*FieldMeta
}

type FieldMeta struct {
	Name    string        `json:"name"`
	Kind    string        `json:"kind"`
	Type    string        `json:"type"`
	Imports []*ImportMeta `json:"imports"`
	Raws    []*CodeMeta   `json:"raws"`
}

type ImportMeta struct {
	Path  string `json:"path"`
	Alias string `json:"alias"`
}

type CodeMeta struct {
	Name string `json:"name"`
	From string `json:"from"`
	Code string `json:"code"`
}
