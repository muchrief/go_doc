package go_doc

type Api interface {
	GetFullUriPath() string
	GetTags() []string
	GetSummary() string
	GetDescription() string
	GetKeyName() string
	GetQuery() any
	GetPayload() any
	GetResponse() any
	GetMethod() string
	GetGroupPath() string
	GetRelativePath() string

	SetGroupPath(path string)
}

type TemplateFunc func() (string, error)
type HandlerDocumentation func(method, path string, template TemplateFunc)

type GoDocumentation interface {
	SetInfo(info *Info) GoDocumentation
	AddServer(server *ServerObject) GoDocumentation
	AddSecurity(key SecuritySchemaKey, security *SecuritySchemeObject) GoDocumentation

	RegisterDoc(api Api)
	RegisterDataDocumentation(url string, handler func(method, path string))
	RegisterDocumentation(docType, dataUri, docUri string, handler HandlerDocumentation)
}
