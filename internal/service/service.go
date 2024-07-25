package service

func GetUser(id int) string {
	var result string

	switch id {
	case 1:
		result = "Выся Пупкин"
	case 2:
		result = "Петя Иванов"
	default:
		result = "Пользователь не найден"
	}
	return result
}
