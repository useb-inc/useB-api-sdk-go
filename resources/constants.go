package resources

const (
	// API_BASE_HOST는 api3 서비스의 기본 호스트
	API_BASE_HOST string = "https://api3.useb.co.kr"
	// AUTH_BASE_HOST는 auth 서비스의 기본 호스트
	AUTH_BASE_HOST string = "https://auth.useb.co.kr"
	// BANKING_BASE_HOST는 openapi 서비스의 기본 호스트
	BANKING_BASE_HOST string = "https://openapi.useb.co.kr"

	HeaderKeyContentType   string = "Content-Type"
	HeaderKeyUserAgent     string = "User-Agent"
	HeaderKeyAuthorization string = "Authorization"

	HeaderValueContentTypeJson string = "application/json"
	HeaderValueUserAgent       string = "useb-api-sdk-go/0.1.0"

	DateLayoutISO string = "2006-01-02 15:04:05"
)
