package responses

type RegisterOKResponse struct {
	Response string `example:"user created"`
}

type RegisterConflictResponse struct {
	Response string `example:"user with this phone already exists"`
}
