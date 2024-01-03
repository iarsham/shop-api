package responses

type Success struct {
	Response string `example:"Success"`
}

type InterServerErrorResponse struct {
	Response string `example:"Internal server error"`
}

type PermissionAdminAllowedResponse struct {
	Response string `example:"permission not allowed, just admin user can perform this action"`
}

type PermissionNotAllowedResponse struct {
	Response string `example:"permission not allowed, you can not perform this action"`
}

type DeleteRecordResponse struct{}
