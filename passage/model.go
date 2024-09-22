package passage

import "time"

type Passage struct {
	UID     string
	Plate   string
	Created time.Time
}

func NewPassage(plate string) *Passage {
	passage := &Passage{}

	passage.Plate = plate
	passage.Created = time.Now()

	return passage
}

// Commands
type SavePassageCommand struct {
	Plate   string
	Created time.Time
}

// Queries
type GetPassageQuery struct {
	UID   string
	Plate *string
}

// DTOs
type SavePassageDTO struct {
	Passage Passage
	Image   string
}
