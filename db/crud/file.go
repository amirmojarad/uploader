package crud

import (
	"errors"
	"uploader/ent"
)

func (crud Crud) CreateFile(file *ent.FileEntity, user *ent.User) (*ent.FileEntity, error) {
	return crud.Client.FileEntity.Create().
		SetName(file.Name).
		SetSize(file.Size).
		SetType(file.Type).
		SetOwner(user).
		SetURL(file.URL).
		Save(*crud.Ctx)
}

func (crud Crud) GetAllFiles(username string) ([]*ent.FileEntity, error) {
	user, err := crud.GetUserWithUsername(username)
	if err != nil {
		return nil, err
	}
	return user.QueryFiles().All(*crud.Ctx)
}

func (crud Crud) DeleteFileWithID(username string, id int) (*ent.FileEntity, error) {
	deletedFile, err := crud.Client.FileEntity.Get(*crud.Ctx, id)
	if deletedFile == nil {
		return nil, errors.New("file entity not found")
	}
	user, err := crud.GetUserWithUsername(username)
	if err != nil {
		return nil, err
	}
	_, err = user.Update().RemoveFiles(deletedFile).Save(*crud.Ctx)
	if err != nil {
		return nil, err
	}
	err = crud.Client.FileEntity.DeleteOne(deletedFile).Exec(*crud.Ctx)
	if err != nil {
		return nil, err
	}
	return deletedFile, nil
}
