package app

import (
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var counter = len(redirect)

// (#8) TODO: Get this from a database!
var redirect = map[string]string{
	"google":         "http://www.google.com ",
	"youtube":        "http://www.youtube.com ",
	"facebook":       "http://www.facebook.com ",
	"instagram":      "http://www.instagram.com ",
	"chatgpt":        "http://www.chatgpt.com ",
	"reddit":         "http://www.reddit.com ",
	"wikipedia":      "http://www.wikipedia.org ",
	"bing":           "http://www.bing.com ",
	"x":              "http://www.x.com ",
	"duckduckgo":     "http://www.duckduckgo.com ",
	"amazon":         "http://www.amazon.com ",
	"yahoo":          "http://www.yahoo.com ",
	"linkedin":       "http://www.linkedin.com ",
	"netflix":        "http://www.netflix.com ",
	"msn":            "http://www.msn.com ",
	"office":         "http://www.office.com ",
	"weather":        "http://www.weather.com ",
	"live":           "http://www.live.com ",
	"spotify":        "http://www.spotify.com ",
	"outlook":        "http://www.outlook.com ",
	"ebay":           "http://www.ebay.com ",
	"pinterest":      "http://www.pinterest.com ",
	"imdb":           "http://www.imdb.com ",
	"paypal":         "http://www.paypal.com ",
	"stackoverflow":  "http://www.stackoverflow.com ",
	"cnn":            "http://www.cnn.com ",
	"bbc":            "http://www.bbc.com ",
	"tripadvisor":    "http://www.tripadvisor.com ",
	"booking":        "http://www.booking.com ",
	"indeed":         "http://www.indeed.com ",
	"quora":          "http://www.quora.com ",
	"microsoft":      "http://www.microsoft.com ",
	"apple":          "http://www.apple.com ",
	"news":           "http://www.news.yahoo.com ",
	"googleuk":       "http://www.google.co.uk ",
	"googleau":       "http://www.google.com.au ",
	"googleca":       "http://www.google.com.ca ",
	"googlenz":       "http://www.google.co.nz ",
	"msnuk":          "http://www.msn.co.uk ",
	"redditall":      "http://www.reddit.com/r/all ",
	"stackexchange":  "http://www.stackexchange.com ",
	"wordpress":      "http://www.wordpress.com ",
	"wix":            "http://www.wix.com ",
	"canva":          "http://www.canva.com ",
	"slack":          "http://www.slack.com ",
	"zoom":           "http://www.zoom.us ",
	"jira":           "http://www.jira.com ",
	"trello":         "http://www.trello.com ",
	"medium":         "http://www.medium.com ",
	"hulu":           "http://www.hulu.com ",
	"twitch":         "http://www.twitch.tv ",
	"vice":           "http://www.vice.com ",
	"theguardian":    "http://www.theguardian.com ",
	"bloomberg":      "http://www.bloomberg.com ",
	"forbes":         "http://www.forbes.com ",
	"mashable":       "http://www.mashable.com ",
	"wired":          "http://www.wired.com ",
	"nytimes":        "http://www.nytimes.com ",
	"washingtonpost": "http://www.washingtonpost.com ",
	"techcrunch":     "http://www.techcrunch.com ",
	"engadget":       "http://www.engadget.com ",
}

type Request struct {
	Url string `json:"url"`
}

type ResponseJson struct {
	Short string `json:"short"`
}

// 500 chances to generate a random string that is not in the existing pool
func randomString(counter int, initalString *strings.Builder) {
	// /62 recursively
	// log_message(strconv.Itoa(counter) + " is the current count")
	if counter >= 62 {
		randomString(counter/62, initalString)
	}
	initalString.WriteByte(chars[counter%62])
}

// Function to return the redirected URL if it exists, else return Hello and the unknown key
func Redirect_url(request string, w http.ResponseWriter, r *http.Request) {
	// map.Keys returns an iterator, is being passed to something that uses it to convert iterator to a list, sorted should give some consistency...
	redirect_keys := slices.Sorted(maps.Keys(redirect))
	// log.Print("Redirect keys are: ", redirect_keys, "request is: ", request)

	if slices.Contains(redirect_keys, request) { // If there is something stored
		Log_message("Redirecting to: " + redirect[request])
		// r.Header.Add("Status", "301")
		http.Redirect(w, r, redirect[request], http.StatusTemporaryRedirect)
	} else {
		Log_message("unable to redirect to:" + request)
	}
}

func SendResponse(w http.ResponseWriter, shortened string) {
	var newShortResponse ResponseJson
	newShortResponse.Short = "http://localhost:8080/r/" + shortened

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newShortResponse)
}

// Add the url to the list of known key / value pairs of shorteners to websites
func Add_url(w http.ResponseWriter, website string) {
	// Check if already added - if so return the key value and present
	for key, knownWebsites := range redirect {
		if knownWebsites == website {
			Log_message("Already added as: " + key)
			SendResponse(w, key)
			return
		}
	}
	// base 62 endocde it - giving me billions of URLs with tiny amounts characters after the url
	var randText strings.Builder
	randomString(counter, &randText)

	// redirect[randText.String()] = "https://" + strings.Split(r.URL.String(), "/")[2]
	redirect[randText.String()] = website
	log.Print(randText.String(), " to go to website: ", website)

	SendResponse(w, randText.String())
	counter += 1
}

var StartPage = func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./template/index.html")
}

var RedirectPage = func(w http.ResponseWriter, r *http.Request) {
	// [1:] to remove the first /
	var request string = strings.Split(r.URL.String()[1:], "r/")[1]
	Log_message("request is: " + request)
	// wrapped in a function so the contents can change without breaking anything else i want to do here
	Redirect_url(request, w, r)
}

var AddURL = func(w http.ResponseWriter, r *http.Request) {
	var request Request
	ErrDecodingJson := json.NewDecoder(r.Body).Decode(&request)
	if ErrDecodingJson != nil {

		expectedJson, _ := json.Marshal(request)
		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorMessage: ErrDecodingJson.Error(),
			UserMessage:  fmt.Sprintf("Error decoding json - this was unable to be decoded as valid json. \n Expexted %s", string(expectedJson)),
		})
		Log_message("[ERROR]: " + ErrDecodingJson.Error())
		return
	}

	var website string = request.Url

	// CHECK THE STRING IS A URL on the back end as well
	_, errParsingURL := url.ParseRequestURI(website)
	if errParsingURL != nil {

		json.NewEncoder(w).Encode(ErrorResponse{
			ErrorMessage: errParsingURL.Error(),
			UserMessage:  "The url you requested was not recognised as a url",
		})
		Log_message("[ERROR]: " + errParsingURL.Error())
		return
	}

	Log_message("Website requested is: " + website)
	if len(request.Url) > 0 { // Check to see if there is a website after add and it's got the right amount
		// wrapped in a function so the contents can change without breaking anything else i want to do here
		Add_url(w, website)
	}
}
