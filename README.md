# go_doc
```
r := gin.Default()

sdk := gin_api.NewGinApiSdk(r)
doc := SetUpDocumentation(r)

sdk.Use(func(a gin_api.Api) {
	doc.RegisterDoc(a)
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

sdk.Register(&gin_api.ApiData{
	Method:       http.MethodGet,
	RelativePath: "/health",
	Response:     &BaseResponse[interface{}]{},
	Tags:         []string{"status"},
}, func(ctx *gin.Context) {
	response := &BaseResponse[interface{}]{
		Message: "OK",
	}
	ctx.JSON(http.StatusOK, response)
})
```