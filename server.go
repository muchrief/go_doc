package go_doc

func NewServer(url, desc string) *ServerObject {
	return &ServerObject{
		Url:         url,
		Description: desc,
	}
}

type serverVariable struct {
	Enum        []string `yaml:"enum" json:"enum"`
	Default     string   `yaml:"default" json:"default"`
	Description string   `yaml:"description" json:"description"`
}

type ServerObject struct {
	Url         string                     `yaml:"url" json:"url"`
	Description string                     `yaml:"description" json:"description"`
	Variables   map[string]*serverVariable `yaml:"variables,omitempty" json:"variables,omitempty"`
}

func (s *ServerObject) SetVariables(key, desc, defaultValue string, enums []string) *ServerObject {
	if s.Variables == nil {
		s.Variables = map[string]*serverVariable{}
	}

	variable := serverVariable{
		Description: desc,
		Enum:        enums,
		Default:     defaultValue,
	}

	s.Variables[key] = &variable
	return s
}
