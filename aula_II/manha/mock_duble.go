package main

type MockRepository struct {
	SearchByNameWasCalled bool
}

// A mistura do conceito de stub e spy, é o mock
func (m *MockRepository) SearchByName(nombre string) string {
	// Ao criarmos uma variável de controle para verificar se o método foi chamado
	// estamos utilizando o conceito de spy
	m.SearchByNameWasCalled = true

	// ao retornarmos um valor hard coded, estamos usando o conceito de "stub"
	return "12345678"
}

func (m *MockRepository) SearchByPhone(phone string) string {
	return ""
}

func (m *MockRepository) AddEntry(name, phone string) error {
	return nil
}
