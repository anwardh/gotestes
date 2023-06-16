package main

type SpyRepository struct {
	SearchByPhoneWasCalled bool
}

func (s *SpyRepository) SearchByName(name string) string {
	return ""
}

func (s *SpyRepository) SearchByPhone(phone string) string {
	s.SearchByPhoneWasCalled = true
	return ""
}

func (s *SpyRepository) AddEntry(name, phone string) error {
	return nil
}
