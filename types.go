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
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	TypeMetas []*TypeMeta `json:"typeMetas"`
}

type TypeMeta struct {
	Id     string      `json:"id"`
	Name   string      `json:"name"`
	From   string      `json:"from"`
	Code   string      `json:"code"`
	Import *ImportMeta `json:"imports"`
}

type ImportMeta struct {
	Path  string `json:"path"`
	Alias string `json:"alias"`
}
