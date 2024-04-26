package apigatewayv2

import "github.com/Drafteame/engine/internal/response"

// HTTPRequest contains data coming from the new HTTP API Gateway
type HTTPRequest struct {
	Version               string             `json:"version"`
	RouteKey              string             `json:"routeKey"`
	RawPath               string             `json:"rawPath"`
	RawQueryString        string             `json:"rawQueryString"`
	Cookies               []string           `json:"cookies,omitempty"`
	Headers               map[string]string  `json:"headers"`
	QueryStringParameters map[string]string  `json:"queryStringParameters,omitempty"`
	PathParameters        map[string]string  `json:"pathParameters,omitempty"`
	RequestContext        HTTPRequestContext `json:"requestContext"`
	StageVariables        map[string]string  `json:"stageVariables,omitempty"`
	Body                  string             `json:"body,omitempty"`
	IsBase64Encoded       bool               `json:"isBase64Encoded"`
}

// HTTPRequestContext contains the information to identify the AWS account and resources invoking the Lambda function.
type HTTPRequestContext struct {
	RouteKey       string                                   `json:"routeKey"`
	AccountID      string                                   `json:"accountId"`
	Stage          string                                   `json:"stage"`
	RequestID      string                                   `json:"requestId"`
	Authorizer     *HTTPRequestContextAuthorizerDescription `json:"authorizer,omitempty"`
	APIID          string                                   `json:"apiId"` // The API Gateway HTTP API Id
	DomainName     string                                   `json:"domainName"`
	DomainPrefix   string                                   `json:"domainPrefix"`
	Time           string                                   `json:"time"`
	TimeEpoch      int64                                    `json:"timeEpoch"`
	HTTP           HTTPRequestContextHTTPDescription        `json:"http"`
	Authentication HTTPRequestContextAuthentication         `json:"authentication,omitempty"`
}

// HTTPRequestContextAuthentication contains authentication context information for the request caller including client certificate information if using mTLS.
type HTTPRequestContextAuthentication struct {
	ClientCert HTTPRequestContextAuthenticationClientCert `json:"clientCert"`
}

// HTTPRequestContextAuthenticationClientCert contains client certificate information for the request caller if using mTLS.
type HTTPRequestContextAuthenticationClientCert struct {
	ClientCertPem string                                             `json:"clientCertPem"`
	IssuerDN      string                                             `json:"issuerDN"`
	SerialNumber  string                                             `json:"serialNumber"`
	SubjectDN     string                                             `json:"subjectDN"`
	Validity      HTTPRequestContextAuthenticationClientCertValidity `json:"validity"`
}

// HTTPRequestContextAuthenticationClientCertValidity contains client certificate validity information for the request caller if using mTLS.
type HTTPRequestContextAuthenticationClientCertValidity struct {
	NotAfter  string `json:"notAfter"`
	NotBefore string `json:"notBefore"`
}

// HTTPRequestContextAuthorizerDescription contains authorizer information for the request context.
type HTTPRequestContextAuthorizerDescription struct {
	JWT    *HTTPRequestContextAuthorizerJWTDescription `json:"jwt,omitempty"`
	Lambda map[string]interface{}                      `json:"lambda,omitempty"`
	IAM    *HTTPRequestContextAuthorizerIAMDescription `json:"iam,omitempty"`
}

// HTTPRequestContextAuthorizerJWTDescription contains JWT authorizer information for the request context.
type HTTPRequestContextAuthorizerJWTDescription struct {
	Claims map[string]string `json:"claims"`
	Scopes []string          `json:"scopes,omitempty"`
}

// HTTPRequestContextAuthorizerIAMDescription contains IAM information for the request context.
type HTTPRequestContextAuthorizerIAMDescription struct {
	AccessKey       string                                      `json:"accessKey"`
	AccountID       string                                      `json:"accountId"`
	CallerID        string                                      `json:"callerId"`
	CognitoIdentity HTTPRequestContextAuthorizerCognitoIdentity `json:"cognitoIdentity,omitempty"`
	PrincipalOrgID  string                                      `json:"principalOrgId"`
	UserARN         string                                      `json:"userArn"`
	UserID          string                                      `json:"userId"`
}

// HTTPRequestContextAuthorizerCognitoIdentity contains Cognito identity information for the request context.
type HTTPRequestContextAuthorizerCognitoIdentity struct {
	AMR            []string `json:"amr"`
	IdentityID     string   `json:"identityId"`
	IdentityPoolID string   `json:"identityPoolId"`
}

// HTTPRequestContextHTTPDescription contains HTTP information for the request context.
type HTTPRequestContextHTTPDescription struct {
	Method    string `json:"method"`
	Path      string `json:"path"`
	Protocol  string `json:"protocol"`
	SourceIP  string `json:"sourceIp"`
	UserAgent string `json:"userAgent"`
}

// HTTPResponse configures the response to be returned by API Gateway V2 for the request
type HTTPResponse struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
	Cookies           []string            `json:"cookies"`
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

func (r *HTTPResponse) SetCookies(cookies []string) {
	r.Cookies = cookies
}

var _ response.Out = (*HTTPResponse)(nil)
