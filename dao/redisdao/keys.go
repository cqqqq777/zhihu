package redisdao

func GetVerificationKey(email string) string {
	return email + ":verification"
}
