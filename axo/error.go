package axo

// Unwrap panics if err is not nil, otherwise returns v
func Unwrap[T any](v T, err error) T {
	if err != nil {
		panic(err) // Rust'taki unwrap() gibi hata alırsa programı durdurur
	}
	return v
}

// Ok returns a pointer to the value if there is no error, otherwise nil
func Ok[T any](v T, err error) *T {
	if err != nil {

		println(err.Error())
		return nil
	}
	return &v
}

// HowTF : how the f*** did this happen *thats a joke don't use this :))*
func HowTF(err error) {
	if err != nil {
		println("💀 How the f*** did this happen?")
		panic(err)
	}
}
