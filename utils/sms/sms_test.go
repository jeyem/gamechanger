package sms

import "testing"

func TestSensSMS(t *testing.T) {
	s := New("6A34415A794E6D456C357047746856397335435353513D3D")
	if err := s.Send("09369323339", "this is test"); err != nil {
		t.Error(err)
	}
}
