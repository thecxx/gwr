package gwr

type remote struct {
	wr *Gwr
}

func (r *remote) IsMobile() bool {
	return false
}
