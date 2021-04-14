package main

// START OMIT

type FS interface {
	Open(name string) (File, error)
}
type File interface {
	//Stat() (FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}
// END OMIT
