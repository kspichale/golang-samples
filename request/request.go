package request

type httpMethod string

// HTTP Methods
const (
	GET    httpMethod = "get"
	PUT               = "put"
	POST              = "post"
	DELETE            = "delete"
)

type request interface {
	Execute() string
}

type requestBuilder interface {
	Method(httpMethod) requestBuilder
	Uri(string) requestBuilder
	Build() request
}

type requestData struct {
	method httpMethod
	uri    string
}

type requestBuilderData struct {
	method httpMethod
	uri    string
}

func New() requestBuilder {
	return &requestBuilderData{}
}

func Request() requestBuilder {
	return &requestBuilderData{}
}

func (builder *requestBuilderData) Method(method httpMethod) requestBuilder {
	builder.method = method
	return builder
}

func (builder *requestBuilderData) Uri(uri string) requestBuilder {
	builder.uri = uri
	return builder
}

func (builder *requestBuilderData) Build() request {
	return &requestData{
		method: builder.method,
		uri:    builder.uri,
	}
}

func (builder *requestData) Execute() string {
	return "result"
}
