package data

var dfEncoder = DefaultEncoder{}

type Encoder interface {
	Encode(source []byte) []byte
}

type DefaultEncoder struct{}

func (e DefaultEncoder) Encode(source []byte) []byte {
	return source
}

func Encode(encoder Encoder, source string) string {
	bytes := []byte(source)
	if encoder != nil {
		return string(encoder.Encode(bytes))
	}
	return string(GetDefaultEncoder().Encode(bytes))
}

func GetDefaultEncoder() Encoder {
	return dfEncoder
}
