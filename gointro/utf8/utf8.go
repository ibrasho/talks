package utf8

type رقم int
type أرقام []int

func (أ أرقام) مجموع() رقم {
	var مجموع رقم = 0

	for س := range أ {
		مجموع = مجموع + رقم(س)
	}

	return رقم(مجموع)
}

func (أ أرقام) مجموع2() int {
	مجموع := 0

	for س := range أ {
		مجموع = مجموع + س
	}

	return مجموع
}
