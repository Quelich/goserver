package main

// Build a path to html file
func buildHtml(path string) string {
	p := STATIC + path + ".html"
	return p
}
