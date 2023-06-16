package main

type StubRepository struct{}

func (d StubRepository) SearchByName(name string) string {
	return "12345678" // note que ao invés de retornar "", retornamos um valor concreto
}

func (d StubRepository) SearchByPhone(phone string) string {
	return ""
}

func (d StubRepository) AddEntry(name, phone string) error {
	return nil
}
