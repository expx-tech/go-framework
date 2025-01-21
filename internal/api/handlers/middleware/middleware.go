package middleware

type loggerHTTP interface {
	DebugHTTP(url, method string)
}

type logger interface {
	Debug(msg string)
}
