package set

type Set map[int]struct{}

func (s Set) Put(x int) {
	s[x] = struct{}{}
}
func (s Set) Has(x int) (exists bool) {
	_, exists = s[x]
	return
}
func (s Set) Remove(val int) {
	delete(s, val)
}
