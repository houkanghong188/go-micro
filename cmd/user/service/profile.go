package service

type UserProfile struct {

}

type User struct {
	ProfileText UserProfile
}
type SelectUserByID func(id string) User

type GetUserProfile func(id string) UserProfile

func NewGetUserProfile(selectUser SelectUserByID) GetUserProfile {
	return func(id string) string {
		user := selectUser(id)
		return user.ProfileText
	}
}

func selectUser(id string) User {
	return User{}
}
