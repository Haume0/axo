package axo

/*
	Quality of Life Functions
	These functions are designed to make error handling easier and preventing code duplication.
	They are not meant to be used in certain cases, such as no way to get the error or when
	the error is not important.
*/

// Unwrap is a utility function that returns the value if there is no error, otherwise it panics.
func Unwrap[T any](v T, err error) T {
	if err != nil {
		panic(err) // Rust'taki unwrap() gibi hata alırsa programı durdurur
	}
	return v
}

// Ok is a utility function that returns a pointer to the value if there is no error, otherwise it prints the error and returns nil.
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
