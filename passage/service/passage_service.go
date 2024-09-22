package service

import (
	"golang_grafana_struct/passage"
)

var (
	_ passage.PassageService = (*PassageServiceImpl)(nil)
)

type PassageServiceImpl struct{}

func ProvidePassageServiceImpl() *PassageServiceImpl {
	passageSvc := &PassageServiceImpl{}
	return passageSvc
}

func (s *PassageServiceImpl) GetPassage(query passage.GetPassageQuery) int {
	return 1
}

func (s *PassageServiceImpl) SavePassage(dto passage.SavePassageDTO) int {
	return 1
}
