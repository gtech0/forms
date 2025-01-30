package get

import "mime/multipart"

type FileDto struct {
	File *multipart.FileHeader `form:"file"`
}
