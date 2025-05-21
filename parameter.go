package go_doc

func NewListParameterObject(data interface{}) ListParameterObject {
	results := ListParameterObject{}

	return results
}

type ParameterObject struct {
	Name        string  `yaml:"name" json:"name"`
	In          string  `yaml:"in" json:"in"`
	Description string  `yaml:"description,omitempty" json:"description,omitempty"`
	Required    bool    `yaml:"required,omitempty" json:"required,omitempty"`
	Deprecated  bool    `yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
	Schema      *Schema `yaml:"schema,omitempty" json:"schema,omitempty"`
}

type ListParameterObject []*ParameterObject
