// select random value from source

const JARGONFILE = "jargonfile"

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
    jargonFile := JargonFile{Jargons: make(map[string]JargonSource), filePath: filePath}
    jargFile := path.Join(filePath, JARGONFILE)
}

