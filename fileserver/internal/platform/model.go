package platform

//FileInfo contains info, given by file miner
type FileInfo struct {
	ID      int    `json:"-" xml:"-"`
	Name    string `json:"name" xml:"name" vls:"required"`
	Mode    string `json:"mode" xml:"mode" vls:"required"`
	ModTime string `json:"mod_time" xml:"mod_time" vls:"required"`
	Size    int64  `json:"size" xml:"size" vls:"required"`
}
