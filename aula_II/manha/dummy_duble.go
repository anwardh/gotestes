package main

type DummyRepository struct{}

func (d DummyRepository) SearchByName(name string) string {
	return ""
}

func (d DummyRepository) SearchByPhone(phone string) string {
	return ""
}

func (d DummyRepository) AddEntry(name, phone string) error {
	return nil
}
