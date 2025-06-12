package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muchrief/go_doc"
)

func SetUpDocumentation(r *gin.Engine) go_doc.GoDocumentation {
	doc := go_doc.NewGoDocumentation(go_doc.InfoVersion3)
	doc.
		SetInfo(&go_doc.Info{
			Title:          "Simple Golang Documentation OpenApi",
			Summary:        "Test simple documentation golang openapi",
			Description:    "Test simple documentation golang openapi",
			TermsOfService: "https://jhondoe.com",
			License:        go_doc.LicenceApache,
			Version:        "v1.0.0",
			Contact: &go_doc.Contact{
				Name:  "John Doe",
				Email: "jhondoe@gmail.com",
				Url:   "https://jhondoe.com",
			},
		}).
		AddServer(&go_doc.ServerObject{
			// Url:         "http://localhost:3000",
			Description: "Development Server",
		}).
		AddSecurity(go_doc.SecuritySchemaKeyOAuth2, &go_doc.SecuritySchemeObject{
			Type: go_doc.SecurityTypeOauth2,
			Flows: &go_doc.OAuthFlowsObject{
				Password: &go_doc.OAuthFlowObject{
					TokenUrl: "https://jhondoe.com/oauth/token",
				},
			},
		})

	// register data
	r.Handle(http.MethodGet, "/doc_data", func(ctx *gin.Context) {
		doc.RegisterDataDocumentation("/doc_data", func(method, path string) {
			ctx.YAML(200, doc)
		})
	})

	// register page
	doc.RegisterDocumentation(
		"redoc",
		"/doc_data",
		"/redoc",
		func(method, path string, template go_doc.TemplateFunc) {
			r.Handle(method, path, func(ctx *gin.Context) {
				templ, err := template()
				if err != nil {
					response := BaseResponse[interface{}]{}
					ctx.JSON(http.StatusInternalServerError, &response)
					return
				}

				ctx.Data(http.StatusOK, "text/html", []byte(templ))
			})
		})

	doc.RegisterDocumentation(
		"swagger",
		"/doc_data",
		"/docs",
		func(method, path string, template go_doc.TemplateFunc) {
			r.Handle(method, path, func(ctx *gin.Context) {
				templ, err := template()
				if err != nil {
					response := BaseResponse[interface{}]{}
					ctx.JSON(http.StatusInternalServerError, &response)
					return
				}

				ctx.Data(http.StatusOK, "text/html", []byte(templ))
			})
		})

	return doc
}
