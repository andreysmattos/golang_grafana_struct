package service

import (
	"golang_grafana_struct/passage"
)

var (
	_ passage.PassageService = (*PassageServiceImpl)(nil)
)

type PassageServiceImpl struct {
	passageStore passage.Store
}

func ProvidePassageService(passageStore passage.Store) passage.PassageService {
	passageSvc := &PassageServiceImpl{
		passageStore: passageStore,
	}
	return passageSvc
}

func (s *PassageServiceImpl) GetPassage(query passage.GetPassageQuery) int {
	s.passageStore.GetPassageById(query)
	return 1
}

func (s *PassageServiceImpl) SavePassage(dto passage.SavePassageDTO) int {

	cmd := passage.SavePassageCommand{
		Plate:   dto.Passage.Plate,
		Created: dto.Passage.Created,
	}

	s.passageStore.SavePassage(cmd)
	return 1
}
