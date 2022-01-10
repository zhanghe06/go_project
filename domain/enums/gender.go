package enums

type GenderType int

// 性别（1:男,2:女）
const (
	_ = iota
	GenderMale
	GenderFemale
)

var genderMap = map[GenderType]string{
	GenderMale:   "男",
	GenderFemale: "女",
}

func (t GenderType) DisplayName() string {
	return genderMap[t]
}
