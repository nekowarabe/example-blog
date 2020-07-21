package entities

// Gender 性別
type Gender uint8

const (
	// GenderUnknown 未提供性別
	GenderUnknown Gender = iota
	// GenderMale 男性
	GenderMale
	// GenderFemale 女性
	GenderFemale
)

// IsValid 是否為有效的性別值
func (gender Gender) IsValid() bool {
	return GenderUnknown <= gender && gender <= GenderFemale
}

// Account 帳號定義
type Account struct {
	Username string
	Password string
	Tel      string
	Age      uint8
	Gender   Gender
}
