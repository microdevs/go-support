package vault

var _ Client = &MockClient{}

// MockClient mock for a client
type MockClient struct {
	WriteDataFunc func(path string, data map[string]interface{}) error
	ReadDataFunc  func(path string) (map[string]interface{}, error)
}

func (m *MockClient) WriteData(path string, data map[string]interface{}) error {
	return m.WriteDataFunc(path, data)
}

func (m *MockClient) ReadData(path string) (map[string]interface{}, error) {
	return m.ReadDataFunc(path)
}
