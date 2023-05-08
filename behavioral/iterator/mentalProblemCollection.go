package iterator

type MentalProblemCollection struct {
	MentalProblems []*MentalProblem
}

func (mp *MentalProblemCollection) CreateIterator() Iterator {
	return &MentalProblemIterator{
		MentalProblems: mp.MentalProblems,
	}
}
