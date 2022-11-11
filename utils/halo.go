package utils

func GeneratePass(password string, salt string) (encryptPass string) {
	encryptPass, err := GenerateMD5(password + salt)
	if err != nil {
		return ""
	}
	return encryptPass
}
