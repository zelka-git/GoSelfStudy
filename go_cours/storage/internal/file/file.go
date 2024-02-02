package file

import ("github.com/google/uuid"
"fmt")

type File struct{
	Id uuid.UDDI
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error){
	fileId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &File{
		Id: fileId,
		Name: filename,
		Data: blob,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("File(%s, %v)", f.Name, f.ID)
}