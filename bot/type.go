package bot

type baiduNetDiskDate struct {
	Resources []struct {
		Highs struct {
			FilelistFilename []string `json:"filelist.filename"`
			Filename         []string `json:"filename"`
		} `json:"highs"`
		Res struct {
			ID       string `json:"id"`
			Filename string `json:"filename"`
			Size     int64  `json:"size"`
			Isdir    int    `json:"isdir"`
			Pwd      string `json:"pwd"`
			Ctime    string `json:"ctime"`
			Utime    string `json:"utime"`
			Ext      string `json:"ext"`
			Filelist []struct {
				Isdir    int    `json:"isdir"`
				Size     int    `json:"size"`
				Filename string `json:"filename"`
				Ext      string `json:"ext"`
			} `json:"filelist"`
		} `json:"res"`
	} `json:"resources"`
	Total int `json:"total"`
}
