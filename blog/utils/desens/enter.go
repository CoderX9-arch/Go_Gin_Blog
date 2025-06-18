package desens

import "strings"

func DesensitizatiionTel(tel string) string {
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}

func DesensitizatiionEmail(email string) string {
	elist := strings.Split(email, "@")
	if len(elist) != 2 {
		return ""
	}
	return elist[0][:1] + "****" + elist[1]
}
