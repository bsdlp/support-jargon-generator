// select random value from source

type JargonSource struct {
    Source  string `json:"source,omitempty"`
    Jargon  string `json:"jargons,omitempty"`
}

type JargonFile struct {
    Jargons  map[string]JargonSource `json:"jargons,omitempty"`
    filePath string
}

// TODO: load up the source
func LoadSource(filePath string) (*JargonFile, error) {
}

