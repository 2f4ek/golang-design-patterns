package iterator

type MentalProblemIterator struct {
	index          int
	MentalProblems []*MentalProblem
}

func (mp *MentalProblemIterator) HasNext() bool {
	if mp.index < len(mp.MentalProblems) {
		return true
	}
	return false

}
func (mp *MentalProblemIterator) GetNext() *MentalProblem {
	if mp.HasNext() {
		mentalProblem := mp.MentalProblems[mp.index]
		mp.index++
		return mentalProblem
	}
	return nil
}
