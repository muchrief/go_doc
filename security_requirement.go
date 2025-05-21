package go_doc

type SecurityType string

const (
	SecurityTypeApiKey SecurityType = "apiKey"
	SecurityTypeHttp   SecurityType = "http"
	SecurityTypeOauth2 SecurityType = "oauth2"
	SecurityTypeOpenId SecurityType = "openIdConnect"
)

type SecurityIn string

const (
	SecurityInQuery  SecurityIn = "query"
	SecurityInHeader SecurityIn = "header"
	SecurityInCookie SecurityIn = "cookie"
)

type SecuritySchema string

const (
	SecuritySchemaBasic  SecuritySchema = "basic"
	SecuritySchemaApiKey SecurityType   = "apiKey"
	SecuritySchemaBearer SecuritySchema = "bearer"
	SecuritySchemaOAuth2 SecuritySchema = "oauth2"
	SecuritySchemaOpenId SecuritySchema = "openIdConnect"
)

type SecurityBearerFormat string

const (
	SecurityBearerFormatJwt SecurityBearerFormat = "jwt"
)

var securityTypeSchema map[SecurityType]map[SecuritySchema]bool = map[SecurityType]map[SecuritySchema]bool{
	SecurityTypeApiKey: {},
	SecurityTypeHttp: {
		SecuritySchemaBasic:  true,
		SecuritySchemaBearer: true,
	},
	SecurityTypeOauth2: {
		SecuritySchemaOAuth2: true,
	},
	SecurityTypeOpenId: {
		SecuritySchemaOpenId: true,
	},
}

func ValidateSecuritySchema(securityType SecurityType, securitySchema SecuritySchema) bool {
	return securityTypeSchema[securityType][securitySchema]
}

type OAuthFlowObject struct {
	// oauth2 ("implicit", "authorizationCode")
	AuthorizationUrl string `yaml:"authorizationUrl,omitempty" json:"authorizationUrl,omitempty"`

	// oauth2 ("password", "clientCredentials", "authorizationCode")
	TokenUrl string `yaml:"tokenUrl,omitempty" json:"tokenUrl,omitempty"`

	// oauth2
	RefreshUrl string `yaml:"refreshUrl,omitempty" json:"refreshUrl,omitempty"`

	// oauth2
	Scopes map[string]string `yaml:"scopes,omitempty" json:"scopes,omitempty"`
}

type OAuthFlowsObject struct {
	Implicit          *OAuthFlowObject `yaml:"implicit,omitempty" json:"implicit,omitempty"`
	Password          *OAuthFlowObject `yaml:"password,omitempty" json:"password,omitempty"`
	ClientCredentials *OAuthFlowObject `yaml:"clientCredentials,omitempty" json:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlowObject `yaml:"authorizationCode,omitempty" json:"authorizationCode,omitempty"`
}

type SecuritySchemeObject struct {
	Type        SecurityType `yaml:"type" json:"type"`
	Description string       `yaml:"description,omitempty" json:"description,omitempty"`

	Name string     `yaml:"name,omitempty" json:"name,omitempty"` //apiKey
	In   SecurityIn `yaml:"in,omitempty" json:"in,omitempty"`     // apiKey

	Scheme       SecuritySchema       `yaml:"scheme,omitempty" json:"scheme,omitempty"`             // http
	BearerFormat SecurityBearerFormat `yaml:"bearerFormat,omitempty" json:"bearerFormat,omitempty"` // http (bearer)

	Flows *OAuthFlowsObject `yaml:"flows,omitempty" json:"flows,omitempty"` // oauth2

	OpenIdConnectUrl string `yaml:"openIdConnectUrl,omitempty" json:"openIdConnectUrl,omitempty"` // openIdConnect
}

type SecurityRequirementObject []*SecuritySchemeObject
