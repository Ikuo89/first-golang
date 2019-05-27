package domain

import (
  "github.com/projects/app/lib"
  "github.com/projects/app/model"
)

var db = lib.DbConnection()

type UserDomain struct {
  *model.User
}

func (user *UserDomain) Save() {
  db.Save(user.User)
}

func (user *UserDomain) Delete() {
  db.Delete(user.User)
}

func GetUsers() []UserDomain {
  users := []model.User{}
  db.Find(&users)

  user_domains := []UserDomain{}
  for _, user := range users {
    copied := user
    user_domains = append(user_domains, UserDomain{&copied})
  }

  return user_domains
}

func GetUser(id uint64) UserDomain {
  user := model.User{}
  db.Where("ID = ?", id).First(&user)

  return UserDomain{&user}
}
