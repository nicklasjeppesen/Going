package cors

// List of allowed origins

func AllowedOrigins() []string {

	return []string{
		"http://localhost:3000",
		"*",
	}
}
