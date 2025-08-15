package utils

func CekStatusUser(status bool) (bool, string) {
	switch status {
	case false:
		return false, "Tidak Aktif"
	case true:
		return true, "Aktif"
	default:
		return false, "Aktif"
	}
}
