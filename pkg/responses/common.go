package responses

type InterServerErrorResponse struct {
	Response string `example:"Internal server error"`
}

type DeleteRecordResponse struct{}

type PermissionNotAllowedResponse struct {
	Response string `example:"permission not allowed, just admin user can perform this action"`
}
