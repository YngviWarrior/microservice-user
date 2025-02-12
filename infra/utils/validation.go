package utils

import "regexp"

func IsValidEmail(email string) bool {
	// Expressão regular para validar o formato do e-mail
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
