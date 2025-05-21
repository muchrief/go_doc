package go_doc

import "net/http"

func NewPathItemObject(summary, desc string) *PathItemObject {
	return &PathItemObject{
		Summary:     summary,
		Description: desc,
	}
}

type PathItemObject struct {
	Summary     string             `yaml:"summary,omitempty" json:"summary,omitempty"`
	Description string             `yaml:"description,omitempty" json:"description,omitempty"`
	Get         *OperationObject   `yaml:"get,omitempty" json:"get,omitempty"`
	Put         *OperationObject   `yaml:"put,omitempty" json:"put,omitempty"`
	Post        *OperationObject   `yaml:"post,omitempty" json:"post,omitempty"`
	Delete      *OperationObject   `yaml:"delete,omitempty" json:"delete,omitempty"`
	Options     *OperationObject   `yaml:"options,omitempty" json:"options,omitempty"`
	Head        *OperationObject   `yaml:"head,omitempty" json:"head,omitempty"`
	Patch       *OperationObject   `yaml:"patch,omitempty" json:"patch,omitempty"`
	Trace       *OperationObject   `yaml:"trace,omitempty" json:"trace,omitempty"`
	Servers     []*ServerObject    `yaml:"servers,omitempty" json:"servers,omitempty"`
	Parameters  []*ParameterObject `yaml:"parameters,omitempty" json:"parameters,omitempty"`
}

// SetOperationObject implements PathItemObject.
func (p *PathItemObject) SetOperationObject(method string, operation *OperationObject) *PathItemObject {
	switch method {
	case http.MethodGet:
		p.Get = operation
	case http.MethodPost:
		p.Post = operation
	case http.MethodPut:
		p.Put = operation
	case http.MethodDelete:
		p.Delete = operation
	}

	return p
}

// SetParameters implements PathItemObject.
func (p *PathItemObject) SetParameters(data interface{}) *PathItemObject {
	if p.Parameters == nil {
		p.Parameters = ListParameterObject{}
	}

	parameters := NewListParameterObject(data)
	p.Parameters = append(p.Parameters, parameters...)

	return p
}
