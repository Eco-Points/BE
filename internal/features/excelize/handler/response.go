package handler

type Url struct {
	Tanggal string `json:"tanggal"`
	Link    string `json:"link"`
}

func toUrlResponse(tgl string, li string) Url {
	return Url{
		Tanggal: tgl,
		Link:    li,
	}
}
