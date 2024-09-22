package passage

type PassageService interface {
	GetPassage(query GetPassageQuery) int
	SavePassage(dto SavePassageDTO) int
}

type Store interface {
	GetPassageById(query GetPassageQuery) int
	SavePassage(cmd SavePassageCommand) int
}
