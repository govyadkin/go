package unique

type Options struct {
	Count     bool
	Duplicate bool
	Unique    bool
	Fields    bool
	NumFields int
	Chars     bool
	NumChars  int
	Register  bool
}

type Data struct {
	Origin string
	Count  int
}
