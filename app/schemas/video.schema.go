package schemas

type VideoMain struct {
	Host          string `json:"host"`
	Title         string `json:"title"`
	Background    string `json:"background"`
	View          int64  `json:"view"`
	Id            string `json:"id"`
	Status        string `json:"status"`
	Started       int64  `json:"started"`
	Desc          string `json:"desc"`
	NameGalxe     string `json:"name_galxe" bson:"name_galxe"`
	LinkGalxe     string `json:"link_galxe" bson:"link_galxe"`
	ImageGalxeUrl string `json:"image_galxe_url" bson:"image_galxe_url"`
}

type VideoPrevious struct {
	Host       string `json:"host"`
	Title      string `json:"title"`
	Background string `json:"background"`
	View       int64  `json:"view"`
	Id         string `json:"id"`
	End        int64  `json:"end"`
}

type PreviousStream struct {
	Host       string `json:"host"`
	Title      string `json:"title"`
	Background string `json:"background"`
	View       int64  `json:"view"`
	Id         string `json:"id"`
	Started    int64  `json:"started"`
	Desc       string `json:"desc"`
}

type PreviousStreamRecommend struct {
	Host       string `json:"host"`
	Title      string `json:"title"`
	Background string `json:"background"`
	View       int64  `json:"view"`
	Id         string `json:"id"`
	Started    int64  `json:"started"`
}
