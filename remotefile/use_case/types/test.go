package types

type TestType struct {
	T
	*Type
	Uri     string
	Content []byte
}

func (t TestType) GetUri() string {
	return t.Uri
}

func (t TestType) GetProtocols() []string {
	return []string{"test"}
}

func (t TestType) Read() ([]byte, error) {
	err := t.Validate()
	if err != nil {
		return []byte{}, err
	}

	return t.Content, nil
}

func (t TestType) Write(content []byte) error {
	err := t.Validate()
	if err != nil {
		return err
	}
	t.Content = content
	return nil
}

func (t TestType) Validate() error {
	return t.validateUri(t)
}

func (t TestType) Sha256() (string, error) {
	return t.sha256(t)
}

func (t TestType) GetFileName() string {
	return t.getFileName(t)
}

func TestFactory(uri string) (TestType, error) {
	test := TestType{Uri: uri}

	return test, test.Validate()
}