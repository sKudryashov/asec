package platform

//FileInfo contains info, given by file miner
type FileInfo struct {
	ID   int    `json:"-" xml:"-"`
	Name string `json:"name" xml:"name" vls:"required"`
}
