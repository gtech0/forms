package model

type File struct {
	BaseModel
	Bucket string
	Size   int64
}
