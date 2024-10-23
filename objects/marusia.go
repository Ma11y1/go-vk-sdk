package objects

type MarusiaPicture struct {
	ID      int `json:"id"`
	OwnerID int `json:"owner_id"`
}

type MarusiaAudio struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	OwnerID int    `json:"owner_id"`
}

type MarusiaAudioMeta struct {
	Album       string `json:"album"`
	Artist      string `json:"artist"`
	Bitrate     string `json:"bitrate"`
	Duration    string `json:"duration"`
	Genre       string `json:"genre"`
	Kad         string `json:"kad"`
	Md5         string `json:"md5"`
	Md5DataSize string `json:"md5_data_size"`
	Samplerate  string `json:"samplerate"`
	Title       string `json:"title"`
}
