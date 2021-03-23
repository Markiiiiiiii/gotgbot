package model

type Filelist struct {
	Isdir    int    `json:"isdir"`
	Size     int    `json:"size"`
	Filename string `json:"filename"`
	Ext      string `json:"ext"`
}
type Res struct {
	ID       string   `json:"id"`
	Filename string   `json:"filename"`
	Size     int64    `json:"size"`
	Isdir    int      `json:"isdir"`
	Pwd      string   `json:"pwd"`
	Ctime    string   `json:"ctime"`
	Utime    string   `json:"utime"`
	Ext      string   `json:"ext"`
	Flist    Filelist `json:"filelist"`
}
type Highs struct {
	FilelistFilename []string `json:"filelist.filename"`
	Filename         []string `json:"filename"`
}
type Resources struct {
	High Highs `json:"highs"`
	Ress Res   `json:"res"`
}

type NetDiskDate struct {
	Resours []Resources `json:"resources"`
	Total   int         `json:"total"`
}

// type prototype struct {
// 	Resources []struct {
// 		Highs struct {
// 			FilelistFilename []string `json:"filelist.filename"`
// 			Filename         []string `json:"filename"`
// 		} `json:"highs"`
// 		Res struct {
// 			ID       string `json:"id"`
// 			Filename string `json:"filename"`
// 			Size     int64  `json:"size"`
// 			Isdir    int    `json:"isdir"`
// 			Pwd      string `json:"pwd"`
// 			Ctime    string `json:"ctime"`
// 			Utime    string `json:"utime"`
// 			Ext      string `json:"ext"`
// 			Filelist []struct {
// 				Isdir    int    `json:"isdir"`
// 				Size     int    `json:"size"`
// 				Filename string `json:"filename"`
// 				Ext      string `json:"ext"`
// 			} `json:"filelist"`
// 		} `json:"res"`
// 	} `json:"resources"`
// 	Total int `json:"total"`
// }
