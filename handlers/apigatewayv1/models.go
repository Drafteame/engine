package apigatewayv1

import "github.com/Drafteame/engine/internal/response"

// HTTPRequest contains data coming from the API Gateway proxy
type HTTPRequest struct {
	Resource                        string              `json:"resource"` // The resource path defined in API Gateway
	Path                            string              `json:"path"`     // The url path for the caller
	HTTPMethod                      string              `json:"httpMethod"`
	Headers                         map[string]string   `json:"headers"`
	MultiValueHeaders               map[string][]string `json:"multiValueHeaders"`
	QueryStringParameters           map[string]string   `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`
	PathParameters                  map[string]string   `json:"pathParameters"`
	StageVariables                  map[string]string   `json:"stageVariables"`
	RequestContext                  HTTPRequestContext  `json:"requestContext"`
	Body                            string              `json:"body"`
	IsBase64Encoded                 bool                `json:"isBase64Encoded,omitempty"`
}

// HTTPRequestContext contains the information to identify the AWS account and resources invoking the
// Lambda function. It also includes Cognito identity information for the caller.
type HTTPRequestContext struct {
	AccountID         string              `json:"accountId"`
	ResourceID        string              `json:"resourceId"`
	OperationName     string              `json:"operationName,omitempty"`
	Stage             string              `json:"stage"`
	DomainName        string              `json:"domainName"`
	DomainPrefix      string              `json:"domainPrefix"`
	RequestID         string              `json:"requestId"`
	ExtendedRequestID string              `json:"extendedRequestId"`
	Protocol          string              `json:"protocol"`
	Identity          HTTPRequestIdentity `json:"identity"`
	ResourcePath      string              `json:"resourcePath"`
	Path              string              `json:"path"`
	Authorizer        map[string]any      `json:"authorizer"`
	HTTPMethod        string              `json:"httpMethod"`
	RequestTime       string              `json:"requestTime"`
	RequestTimeEpoch  int64               `json:"requestTimeEpoch"`
	APIID             string              `json:"apiId"` // The API Gateway rest API ID
}

// HTTPRequestIdentity contains identity information for the request caller.
type HTTPRequestIdentity struct {
	CognitoIdentityPoolID         string `json:"cognitoIdentityPoolId"`
	AccountID                     string `json:"accountId"`
	CognitoIdentityID             string `json:"cognitoIdentityId"`
	Caller                        string `json:"caller"`
	APIKey                        string `json:"apiKey"`
	APIKeyID                      string `json:"apiKeyId"`
	AccessKey                     string `json:"accessKey"`
	SourceIP                      string `json:"sourceIp"`
	CognitoAuthenticationType     string `json:"cognitoAuthenticationType"`
	CognitoAuthenticationProvider string `json:"cognitoAuthenticationProvider"`
	UserArn                       string `json:"userArn"` //nolint: stylecheck
	UserAgent                     string `json:"userAgent"`
	User                          string `json:"user"`
}

type HTTPResponse struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
}

func (r *HTTPResponse) SetStatusCode(statusCode int) {
	r.StatusCode = statusCode
}

func (r *HTTPResponse) SetHeaders(headers map[string]string) {
	r.Headers = headers
}

func (r *HTTPResponse) SetMultiValueHeaders(mvHeaders map[string][]string) {
	r.MultiValueHeaders = mvHeaders
}

func (r *HTTPResponse) SetBody(body string) {
	r.Body = body
}

func (r *HTTPResponse) SetIsBase64Encoded(b64 bool) {
	r.IsBase64Encoded = b64
}

func (r *HTTPResponse) SetCookies(_ []string) {
	// Not supported in API Gateway V1
}

var _ response.Out = (*HTTPResponse)(nil)
