package models

import (
	"errors"
	"time"
)

var (
	UserList map[string]*User
)


type User struct {
	Id       int64		`json:"id"`
	Username string		`orm:"size(64)"		json:"username"`
	Password string		`orm:"size(128)"		json:"password"`
	Email  string		`orm:"size(128)"		json:"email"`
	Role	string		`orm:"size(64)"		json:"role"`
	AssessCount int		`json:"assessCount"`
}

type IPUser struct {
	Id      int64		`json:"id"`
	Ip 		string		`orm:"size(64)"		json:"ip"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"		json:"createTime"`
	AssessCount int		`json:"assessCount"`
}

func AddUser(u User) int64 {
	// u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	// UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
