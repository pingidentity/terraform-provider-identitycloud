package providererror

const (
	// Common provider error summaries
	InvalidProviderConfigurationError = "Invalid provider configuration"
	AicAPIError                       = "Advanced Identity Cloud API error"
	InternalProviderError             = "Internal provider error"
	DeletedNotRemovedWarning          = "Resource removed from terraform state but not deleted from the service"
)
