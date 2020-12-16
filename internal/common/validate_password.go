package common

var adminPassword string

// ValidatePassword validates a password
func ValidatePassword(password string) bool {
	return password == adminPassword
}

// SetPassword sets the admin password
func SetPassword(password string) {
	adminPassword = password
}