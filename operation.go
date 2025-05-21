package go_doc

func NewOperationObject(tags []string, summary, desc, operationID string) *OperationObject {
	return &OperationObject{
		Tags:        tags,
		Summary:     summary,
		Description: desc,
		OperationId: operationID,
	}
}

type MediaType string

const (
	MediaTypeJson      MediaType = "application/json"
	MediaTypeTextPlain MediaType = "text/plain; charset=utf-8"
)

type MediaTypeObject struct {
	Schema *Schema `yaml:"schema" json:"schema"`
}

type RequestBodyObject struct {
	Description string                         `yaml:"description,omitempty" json:"description,omitempty"`
	Required    bool                           `yaml:"required,omitempty" json:"required,omitempty"`
	Content     map[MediaType]*MediaTypeObject `yaml:"content,omitempty" json:"content,omitempty"`
}

type ResponseObject struct {
	Description string                         `yaml:"description,omitempty" json:"description,omitempty"`
	Headers     map[string]interface{}         `yaml:"headers,omitempty" json:"headers,omitempty"`
	Content     map[MediaType]*MediaTypeObject `yaml:"content,omitempty" json:"content,omitempty"`
	Links       map[string]interface{}         `yaml:"links,omitempty" json:"links,omitempty"`
}

type OperationObject struct {
	Tags        []string                   `yaml:"tags,omitempty" json:"tags,omitempty"`
	Summary     string                     `yaml:"summary,omitempty" json:"summary,omitempty"`
	Description string                     `yaml:"description,omitempty" json:"description,omitempty"`
	OperationId string                     `yaml:"operationId,omitempty" json:"operationId,omitempty"`
	Parameters  ListParameterObject        `yaml:"parameters,omitempty" json:"parameters,omitempty"`
	RequestBody *RequestBodyObject         `yaml:"requestBody,omitempty" json:"requestBody,omitempty"`
	Responses   map[string]*ResponseObject `yaml:"responses,omitempty" json:"responses,omitempty"`
	Callbacks   interface{}                `yaml:"callbacks,omitempty" json:"callbacks,omitempty"`
	Deprecated  bool                       `yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
	Security    interface{}                `yaml:"security,omitempty" json:"security,omitempty"`
	Servers     []*ServerObject            `yaml:"servers,omitempty" json:"servers,omitempty"`
}

// SetParameters implements OperationObject.
func (o *OperationObject) SetParameters(data interface{}) *OperationObject {
	if data == nil {
		return o
	}

	if o.Parameters == nil {
		o.Parameters = ListParameterObject{}
	}

	parameters := NewListParameterObject(data)
	o.Parameters = append(o.Parameters, parameters...)

	return o
}

// SetRequestBody implements OperationObject.
func (o *OperationObject) SetRequestBody(data interface{}) *OperationObject {
	if data == nil {
		return o
	}

	if o.RequestBody == nil {
		o.RequestBody = &RequestBodyObject{}
	}

	schemaPayload := BuildSchema(data)

	o.RequestBody = &RequestBodyObject{
		Description: "",
		Required:    false,
		Content: map[MediaType]*MediaTypeObject{
			MediaTypeJson: {
				Schema: schemaPayload,
			},
		},
	}

	return o
}

// SetResponse implements OperationObject.
func (o *OperationObject) SetResponse(code string, data interface{}) *OperationObject {
	if data == nil {
		return o
	}
	if o.Responses == nil {
		o.Responses = map[string]*ResponseObject{}
	}

	responseSchema := BuildSchema(data)
	o.Responses[code] = &ResponseObject{
		Content: map[MediaType]*MediaTypeObject{
			MediaTypeJson: {
				Schema: responseSchema,
			},
		},
	}

	return o
}
