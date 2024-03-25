package users

type WrongUsernameOrPasswordError struct {
}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

type InvalidUsernameOrPasswordError struct{}

func (m *InvalidUsernameOrPasswordError) Error() string {
	return "username or password must not empty"
}

type InvalidEmailError struct{}

func (m *InvalidEmailError) Error() string {
	return "email must not empty"
}

type InvalidTokenError struct {
	Token string
}

func (m *InvalidTokenError) Error() string {
	return "Invalid token: " + m.Token
}

type InvalidUserError struct {
}

func (m *InvalidUserError) Error() string {
	return "InvalidUserError"
}
