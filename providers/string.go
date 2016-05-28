package providers

type StringProvider struct {
	text string
}

func (this StringProvider) GetBytes() []byte {
	return []byte(this.text)
}

func NewStringProvider(text string) StringProvider {
	return StringProvider{ text }
}
