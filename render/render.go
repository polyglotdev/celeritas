package render

// Render is a struct that contains the configuration for the renderer.
type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
}
