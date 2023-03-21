package mate

type Slot struct{}

type Listen struct{}

type ObjectMate struct {
	Name      string          `json:"name"`
	Kind      string          `json:"kind"`
	Functions []*FunctionMate `json:"functions"`
}

type FunctionMate struct {
	Name       string
	Parameters []*FieldMate
	Results    []*FieldMate
}

type FieldMate struct {
	Name    string        `json:"name"`
	Kind    string        `json:"kind"`
	Type    string        `json:"type"`
	Imports []*ImportMate `json:"imports"`
	Raws    []*CodeMate   `json:"raws"`
}

type ImportMate struct {
	Path  string `json:"path"`
	Alias string `json:"alias"`
}

type CodeMate struct {
	Name string `json:"name"`
	From string `json:"from"`
	Code string `json:"code"`
}
