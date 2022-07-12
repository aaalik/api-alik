package asql

func Decrypt(key, column string, quote int) string {
	if quote == 1 {
		return "CONVERT(AES_DECRYPT(UNHEX('" + column + "'),'" + key + "'),CHAR(255))"
	} else {
		return "CONVERT(AES_DECRYPT(UNHEX(" + column + "),'" + key + "'),CHAR(255))"
	}
}
