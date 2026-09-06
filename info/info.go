package info

type Req struct {
	Label   string `json:"label"`
	Profile string `json:"profile"`
}

type Resp struct {
	Certificate  string   `json:"certificate"`
	Usage        []string `json:"usages"`
	ExpiryString string   `json:"expiry"`
}
