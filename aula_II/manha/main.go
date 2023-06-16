package main

import "fmt"

// A interface da Repository é correspondente a interface "SearchEngine"
type Repository interface {
	SearchByName(name string) string
	SearchByPhone(phone string) string
	AddEntry(name, phone string) error
}

// A struct Service é como se fosse nossa "engine"
type Service struct {
	// o repository é como se fosse nossa "searchEngine"
	repository Repository
}

// type service {
// 	repository Repository
// }

// A funcão que cria uma "Service" que corresponde a nossa "Engine"
func NewService(repository Repository) *Service {
	return &Service{
		repository,
	}
}

func (e *Service) GetVersion() string {
	return "1.0"
}

func (e *Service) FindByName(name string) string {
	return e.repository.SearchByName(name)
}

func (e *Service) SearchByPhone(phone string) string {
	if len(phone) < 5 {
		return ""
	}
	return e.repository.SearchByPhone(phone)
}

func (e *Service) SearchByName(name string) string {
	if len(name) < 5 {
		return ""
	}
	return e.repository.SearchByName(name)
}

func (e *Service) AddEntry(name, phone string) error {
	return e.repository.AddEntry(name, phone)
}

func main() {
	repository := DummyRepository{}
	service := NewService(repository)

	fmt.Printf("A versão atual é %s", service.GetVersion())
}
