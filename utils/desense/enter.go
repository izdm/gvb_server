package desense

import "strings"

func DesensitizationEmail(email string) string {
	// 256655@qq.com 2****@qq.com
	// yaheb7479@yahoo.com y****@yahoo.com
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

func DesensitizationTel(tel string) string {
	//15115790125
	//151****0125
	if len(tel) != 11 {
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
