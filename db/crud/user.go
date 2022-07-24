package crud

import (
	"log"
	"uploader/ent"
	"uploader/ent/user"
	"uploader/usecases"
)

func (crud *Crud) CreateUser(user *ent.User) (*ent.User, error) {
	hashedPassword, err := usecases.HashPassword(user.Password)
	if err != nil {
		log.Printf("error while hashing password: %+v", err)
	}
	return crud.Client.User.Create().SetUsername(user.Username).SetPassword(hashedPassword).Save(*crud.Ctx)
}

func (crud *Crud) GetUserWithUsername(username string) (*ent.User, error) {
	return crud.Client.User.Query().Where(user.UsernameEQ(username)).First(*crud.Ctx)
}
