package filter

func IsAuthServer(host string) bool {

	return EndsWith(host, "okta.com")
}

func IsAnalyticsServer(host string) bool {

	return EndsWith(host, "totango.com", "segment.io", "intercom.io", "pendo.io")
}

func IsFontServer(host string) bool {

	return EndsWith(host, "fonts.googleapis.com", "fonts.gstatic.com")
}

func IsUIStaticFile(path string) bool {

	return EndsWith(path, ".htm", ".", ".jpg", ".jpeg", ".gif", ".png", ".ico", ".css", ".zip", ".tgz", ".gz", ".rar", ".bz2", ".doc", ".xls", ".exe", ".pdf", ".ppt", ".txt", ".tar", ".mid", ".midi", ".wav", ".bmp", ".rtf", ".woff", ".woff2", ".js", ".svg", ".js.map", ".css.map", ".ts.map")
}
