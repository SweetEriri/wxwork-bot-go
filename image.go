package wxworkbot

type imageMessage struct {
	messagetype
	Image Image `json:"image"`
}

type Image struct {
	Base64 string `json:"base64"`
	MD5    string `json:"md5"`
}
