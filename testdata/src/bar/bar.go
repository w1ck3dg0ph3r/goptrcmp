package bar

type Bar struct {
	A string
	B *string
}

func cmp(a, b Bar) bool {
	if a.B == nil {
		return false
	}

	if nil == b.B {
		return false
	}

	ap := a.B
	bp := b.B
	return ap == bp // want "pointer comparison: ap == bp"
}
