package domain

import "strconv"

type User struct {
	ID    string
	Token string
}

type UserDTO struct {
	ID    int64
	Token string
}

func (dto *UserDTO) ToUser() User {
	var u User

	u.Token = dto.Token
	u.ID = strconv.FormatInt(dto.ID, 10)

	return u
}

func (u *User) ToDTO() UserDTO {
	var dto UserDTO

	dto.Token = u.Token
	dto.ID, _ = strconv.ParseInt(u.ID, 10, 64)

	return dto
}