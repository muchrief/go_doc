package go_doc

var BasicSecurityScheme = &SecuritySchemeObject{
	Type:   SecurityTypeHttp,
	Scheme: SecuritySchemaBasic,
}
var BearerSecurityScheme = &SecuritySchemeObject{
	Type:         SecurityTypeHttp,
	Scheme:       SecuritySchemaBearer,
	BearerFormat: SecurityBearerFormatJwt,
}
var ApiKeySecurityScheme = &SecuritySchemeObject{
	Type: SecurityTypeApiKey,
}

type SecuritySchemaKey string

const (
	SecuritySchemaKeyBasic  SecuritySchemaKey = "BasicAuth"
	SecuritySchemaKeyBearer SecuritySchemaKey = "BearerAuth"
	SecuritySchemaKeyApiKey SecuritySchemaKey = "ApiKeyAuth"
	SecuritySchemaKeyOAuth2 SecuritySchemaKey = "OAuth2"
	SecuritySchemaKeyOpenId SecuritySchemaKey = "OpenID"
)

var SecurityTypeSchemaKeyMap = map[SecuritySchema]SecuritySchemaKey{
	SecuritySchemaBasic:  SecuritySchemaKeyBasic,
	SecuritySchemaBearer: SecuritySchemaKeyBearer,
}

type ComponentsObject struct {
	SecuritySchemes map[SecuritySchemaKey]*SecuritySchemeObject `yaml:"securitySchemes,omitempty" json:"securitySchemes,omitempty"`
}

func (c *ComponentsObject) AddSecurityScheme(key SecuritySchemaKey, security *SecuritySchemeObject) *ComponentsObject {
	if c.SecuritySchemes == nil {
		c.SecuritySchemes = map[SecuritySchemaKey]*SecuritySchemeObject{}
	}
	c.SecuritySchemes[key] = security
	return c
}
