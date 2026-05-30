package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const URL ="http://localhost:8000/"

func main(){
	if len(os.Args) < 2  {
		ManPage()
	}
	args := os.Args[1]

	switch args {
	case "health" , "--health":
		HealthChekupGet(URL)
	case "list", "--list":
		GetAllShortLinks(URL)
	case "create", "--create":
		CreateShortLink(URL)
	case "analyse", "--analyse":
		GetAnalytics(URL)
	case "man", "--man":
		ManPage()
	default:
		ManPage()
	}

	
}

func HealthChekupGet(url string){
	res , err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status: ", res.StatusCode)
	fmt.Println("Content length: ", res.ContentLength)
	defer res.Body.Close()
	
	// Method-1
	response, _ := io.ReadAll(res.Body)
	content := string(response)
	fmt.Println("Res: ", content)


	// // Method-2
	// var respString strings.Builder
	// byteCOunt , _ := respString.Write([]byte(content)) 
	// fmt.Println("Count: ",byteCOunt)
	// fmt.Println("Response: ",respString.String())

}

func errHandle(err error){
	if err != nil {
		panic(err)
	}
}

// For /api/v1/url/
type UrlItem struct {
	ID   string    `json:"_id"`
	Name string `json:"createdAt"`
	HashId string `json:"hash_id"`
	Original_url string `json:"original_url"`
	Short_url string `json:"short_url"`
}

type getAllUrl struct{
	Message string `json:"message"`
	Data []UrlItem `json:"data"`
}
func GetAllShortLinks(url string){
	res , err := http.Get((url+"api/v1/url"))

	errHandle(err)

	response , err := io.ReadAll(res.Body)

	defer res.Body.Close()

	var jsonData getAllUrl
	errrs := json.Unmarshal(response, &jsonData)
	
	errHandle(errrs)
	
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------")
	for _,val := range jsonData.Data{
		fmt.Println("Id: ", val.ID)
		fmt.Println("hashId: ", val.HashId)
		fmt.Println("Short url: ", val.Short_url)
		fmt.Println("Original url: ", val.Original_url)
		fmt.Println("-------------------------------------------------------------------------------------------------------------------------")

	}
}


type postUrl struct {
	Original_url string `json:"original_url"`
}
type getUrl struct{
	Message string `json:"message"`
	Data UrlItem `json:"data"`
}

func CreateShortLink(url string){
	scanner  := bufio.NewScanner(os.Stdin)
	var inpUrl string
	fmt.Println("Input a link you want shorten: ")
	if scanner.Scan(){
		inpUrl = scanner.Text()
	}
	payload := postUrl{
		Original_url: inpUrl,
	}
	jsonStr ,err  := json.Marshal(payload)
	
	errHandle(err)

	reqBody := strings.NewReader(string(jsonStr))


	Url := url+"api/v1/url" 
	res , err := http.Post(Url , "application/json", reqBody )
	errHandle(err)

	defer res.Body.Close()


	response , _:= io.ReadAll(res.Body)
	var jsonData getUrl
	stat := json.Unmarshal(response,&jsonData)
	errHandle(stat)

	fmt.Println("---------------------------------------------------------------------------------------------------------------")
	fmt.Println("hash Id: ", jsonData.Data.HashId)
	fmt.Println("hash Id: ", jsonData.Data.Short_url)
	fmt.Println("Original url: ", jsonData.Data.Original_url)
	fmt.Println("---------------------------------------------------------------------------------------------------------------")


}


type analyticsObj struct{
	Id string `json:"_id"`
	Count int `json:"count"`
	IpAddress []string `json:"ipAddress"`
}
type UrlAnalytics struct {
 	Analytics analyticsObj `json:"analytics"`
	Metadata UrlItem `json:"metadata"`

}

func GetAnalytics(url string){
	scanner  := bufio.NewScanner(os.Stdin)
	var inpId string
	fmt.Println("Input the hash ID assigned: ")
	if scanner.Scan(){
		inpId = scanner.Text()
	}


	res ,err := http.Get((url+"analytics/"+inpId))
	errHandle(err)
	
	response , _:= io.ReadAll(res.Body)
	
	defer res.Body.Close()
	var content UrlAnalytics
	stat := json.Unmarshal(response,&content)
	errHandle(stat)
	// fmt.Println("Res:" , content)

	fmt.Println("---------------------------------------------------------------------------------------------------------------")
		fmt.Println("Hash ID: ", content.Metadata.HashId)
		fmt.Println("Clicks: ", content.Analytics.Count)
		fmt.Println("IPs: ", content.Analytics.IpAddress)
		fmt.Println("Original url: ", content.Metadata.Original_url)
		fmt.Println("Short url: ", content.Metadata.Short_url)
	fmt.Println("---------------------------------------------------------------------------------------------------------------")




}

func ManPage(){
	fmt.Println("Hello user this is a small url shortner i made while learning Golang")
	fmt.Print("If you are not sure about the commands please read below docs.\n\n\n")

	fmt.Println("shorten {args}")
	fmt.Println("shorten health or shorten --health")
	fmt.Println("     This check the status of the backend server running on NodeJS")
	fmt.Println("")
	
	fmt.Println("shorten list or shorten --list")
	fmt.Println("     This lists all the url ever listed by you")
	
	fmt.Println("")
	fmt.Println("shorten create or shorten --create")
	fmt.Println("     This creates the short url used to redierect to the provided Url")
	fmt.Println("")
	

	fmt.Println("shorten analyse or shorten --analyse")
	fmt.Println("     This shows analytics on a particular url based on thier hash id provided to you.")
	fmt.Println("")
	

}