package json

func init() {
	if err := parseSchema(); err != nil {
		panic(err)
	}
}
