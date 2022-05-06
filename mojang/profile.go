package mojang

// Profile represents a whole player profile
type Profile struct {
	UUID       string            `json:"id"`
	Name       string            `json:"name"`
	Properties []ProfileProperty `json:"properties"`
}

// ProfileProperty represents a property of a player profile
type ProfileProperty struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature"`
}

// UserInfo represents a property of a user property
type UserInfo struct {
	ID         string             `json:"id"`
	Properties []UserInfoProperty `json:"properties"`
}

// UserInfoProperty represents a property of a user property
type UserInfoProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// MetaInfo authlib-injector meta
type MetaInfo struct {
	Meta               map[string]interface{} `json:"meta"`
	SkinDomains        []string               `json:"skinDomains"`
	SignaturePublicKey string                 `json:"signaturePublickey"`
}
