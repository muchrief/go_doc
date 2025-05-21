package go_doc

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/muchrief/go_doc/view"
)

func NewGoDocumentation(version string) GoDocumentation {
	return &goDocumentationImpl{
		OpenApi: OpenApiVersion(version),
		Info:    NewInfo("Api Documentation", "api documentation", InfoVersion3),
	}
}

type OpenApiVersion string

type goDocumentationImpl struct {
	OpenApi    OpenApiVersion             `yaml:"openapi" json:"openapi"`
	Info       *Info                      `yaml:"info" json:"info"`
	Servers    []*ServerObject            `yaml:"servers,omitempty" json:"servers,omitempty"`
	Paths      map[string]*PathItemObject `yaml:"paths,omitempty" json:"paths,omitempty"`
	Security   SecurityRequirementObject  `yaml:"security,omitempty" json:"security,omitempty"`
	Components *ComponentsObject          `yaml:"components,omitempty" json:"components,omitempty"`
}

// SetInfo implements GoDocumentation.
func (g *goDocumentationImpl) SetInfo(info *Info) GoDocumentation {
	g.Info = info
	return g
}

// AddServer implements GoDocumentation.
func (g *goDocumentationImpl) AddServer(server *ServerObject) GoDocumentation {
	if g.Servers == nil {
		g.Servers = []*ServerObject{}
	}

	g.Servers = append(g.Servers, server)
	return g
}

// AddSecurity implements GoDocumentation.
func (g *goDocumentationImpl) AddSecurity(key SecuritySchemaKey, security *SecuritySchemeObject) GoDocumentation {
	if g.Security == nil {
		g.Security = SecurityRequirementObject{}
	}
	g.Security = append(g.Security, security)

	if g.Components == nil {
		g.Components = &ComponentsObject{}
	}
	g.Components.AddSecurityScheme(key, security)

	return g
}

func (g *goDocumentationImpl) addApi(fullPathUri string, item *PathItemObject) GoDocumentation {
	if g.Paths == nil {
		g.Paths = map[string]*PathItemObject{}
	}

	if !strings.HasPrefix(fullPathUri, "/") {
		fullPathUri = "/" + fullPathUri
	}

	g.Paths[fullPathUri] = item

	return g
}

// Add implements GoDocumentation.
func (g *goDocumentationImpl) RegisterDoc(api Api) {
	operationObject := NewOperationObject(
		api.GetTags(),
		api.GetSummary(),
		api.GetDescription(),
		api.GetKeyName(),
	)
	operationObject.
		SetParameters(api.GetQuery()).
		SetRequestBody(api.GetPayload()).
		SetResponse(
			"200",
			api.GetResponse(),
		)

	pathItem := NewPathItemObject(api.GetSummary(), api.GetDescription())
	pathItem.
		SetOperationObject(
			api.GetMethod(),
			operationObject,
		)

	g.addApi(
		api.GetFullUriPath(),
		pathItem,
	)
}

func (g *goDocumentationImpl) RegisterDataDocumentation(url string, handler func(method, path string)) {
	if url == "" {
		url = "/doc_data"
	}

	handler(http.MethodGet, url)
}

// RegisterDocumentation implements GoDocumentation.
func (g *goDocumentationImpl) RegisterDocumentation(docType, dataUri, docUri string, handler HandlerDocumentation) {
	if dataUri == "" {
		dataUri = "/doc_data"
	}

	if docUri == "" {
		docUri = "/docs"
	}

	var getTemplate func() (string, error)

	switch docType {
	case "swagger":
		getTemplate = func() (string, error) {
			return view.GetSwaggerViewTemplate(&view.ViewTemplateConfig{
				Title: g.Info.GetTitle(),
				Url:   dataUri,
			})
		}
	case "redoc":

		getTemplate = func() (string, error) {
			return view.GetRedocViewTemplate(&view.ViewTemplateConfig{
				Title: g.Info.GetTitle(),
				Url:   dataUri,
			})
		}
	default:
		slog.Error("invalid doc_type")
		return
	}

	handler(http.MethodGet, docUri, getTemplate)
}
