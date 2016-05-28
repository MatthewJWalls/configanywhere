package providers

type StringProvider struct {
	text string
}

func (this StringProvider) GetBytes() ([]byte, error) {
	return []byte(this.text), nil
}

func NewStringProvider(text string) StringProvider {
	return StringProvider{ text }
}
