package main

type FakeRepository struct {
	DB map[string]string
}

// {
// 	"Batata": "Boa",
// 	"Chuchu": "Ruim"
// }

func (m *FakeRepository) SearchByName(name string) string {
	// Se a pessoa passar o nome "Batata" será retornado o valor "Boa"
	return m.DB[name]
}

func (m *FakeRepository) SearchByPhone(phone string) string {
	for nameContact, phoneContact := range m.DB {
		if phoneContact == phone {
			return nameContact
		}
	}
	return ""
}

func (m *FakeRepository) AddEntry(name, phone string) error {
	if m.DB == nil {
		m.DB = make(map[string]string)
	}
	m.DB[name] = phone
	// No caso de uma implementação real de um repository,
	// poderia ocorrer um erro na hora de inserir, retornaríamos esse erro neste momento
	return nil
}
