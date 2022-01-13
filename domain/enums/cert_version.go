package enums

type CertVersion int

// 证书版本（0:V1,1:V2,2:V3）
const (
	V1 = iota
	V2
	V3
)

var certVersionMap = map[CertVersion]string{
	V1: "V1",
	V2: "V2",
	V3: "V3",
}

func (t CertVersion) DisplayName() string {
	return certVersionMap[t]
}
