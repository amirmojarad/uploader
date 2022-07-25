package crud

import "uploader/ent"

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
