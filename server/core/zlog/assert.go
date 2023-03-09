package zlog

func MustNil(err error) {
	if err != nil {
		Fatalf("%v", err)
	}
}
