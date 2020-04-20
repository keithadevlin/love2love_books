//go:generate mockgen -package mws -source=mws.go -destination mws_mocks.go

package mws

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/keithadevlin/love2love_books/pkg/shared/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

type AmazonMWSAPI struct {
	AccessKey       string
	SecretKey       string
	Host            string
	AuthToken       string
	MarketplaceIdUK string
	MarketplaceIdDE string
	SellerId        string
}

type MWSAPI interface {
	genSignAndFetch(Action string, ActionPath string, Parameters map[string]string) (string, error)
	GetLowestOfferListingsForASIN(items []string) (string, error)
	GetLowestOfferListingsForSKU(items []string) (string, error)
	GetCompetitivePricingForASIN(items []string) (string, error)
	GetLowestPricedOffersForSKU(sku string) (string, error)
	GetMatchingProductForId(idType string, idList []string) (string, error)
	SubmitFeed(feedType, feed string) (string, error)
	RequestReport(reportType string) (string, error)
	GetReportRequestList(reportRequestId string) (string, error)
	GetReport(generatedReportId string) (string, error)
}

func NewAmazonMWSAPI(accessKey, secretKey, host, authToken, marketplaceIdUK, marketplaceIdDE, sellerId string) AmazonMWSAPI {
	return AmazonMWSAPI{
		AccessKey:       accessKey,
		SecretKey:       secretKey,
		Host:            host,
		AuthToken:       authToken,
		MarketplaceIdUK: marketplaceIdUK,
		MarketplaceIdDE: marketplaceIdDE,
		SellerId:        sellerId,
	}
}

func (api AmazonMWSAPI) genSignAndFetch(Action string, ActionPath string, Parameters map[string]string) (string, error) {
	genUrl, err := GenerateAmazonUrl(api, Action, ActionPath, Parameters)
	if err != nil {
		return "", err
	}

	SetTimestamp(genUrl)

	signedurl, err := SignAmazonUrl(genUrl, api, Action)
	if err != nil {
		return "", err
	}

	if Action == "SubmitFeed" {

		//feed2 := fmt.Sprintf("sku	price\n1000---0734	9.99")
		//sendbody := []byte(feed2)

		//not needed anymore

		//outputFileName := "/tmp/newPrices.txt"

		var outputFileName string
		if os.Args[1] == "DE" {
			outputFileName = "/tmp/newPricesDE.txt"
		} else {
			outputFileName = "/tmp/newPrices.txt"
		}

		sendbody, err := ioutil.ReadFile(outputFileName)

		//values.Add("Content-MD5", fmt.Sprintf("%x", md5.Sum(data)))
		//md5Value := fmt.Sprintf("%x", md5.Sum(sendbody))

		//_ := http.Header.Add{"Content-MD5": md5Value}
		//

		// below is code we may need to rewrite

		// hooray

		//client := &http.Client{}
		//req, err := http.NewRequest("POST",signedurl, bytes.NewReader(sendbody))
		//req.Header.Add("Content-MD5", md5Value)
		//req.Header.Add("Content-Type", "text/xml")
		//
		//resp, err := client.Do(req)
		//defer resp.Body.Close()

		resp, err := http.Post(signedurl, "text/xml", bytes.NewBuffer(sendbody))
		//http.Header.Add("Content-MD5", md5Value)

		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Unable to get signed URL")
			return "", err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Unable to get read resp body")
			return "", err
		}

		return string(body), nil
	} else if Action == "RequestReport" {
		resp, err := http.PostForm(signedurl, nil)
		//http.Header.Add("Content-MD5", md5Value)

		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Unable to get signed URL")
			return "", err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Unable to get read resp body")
			return "", err
		}

		return string(body), nil
	} else if Action == "GetReportRequestList" {
		resp, err := http.PostForm(signedurl, nil)
		//http.Header.Add("Content-MD5", md5Value)

		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Unable to get signed URL")
			return "", err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Unable to get read resp body")
			return "", err
		}

		return string(body), nil
	} else if Action == "GetReport" {
		resp, err := http.PostForm(signedurl, nil)
		//http.Header.Add("Content-MD5", md5Value)

		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Unable to get signed URL")
			return "", err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Unable to get read resp body")
			return "", err
		}

		return string(body), nil
	} else {
		resp, err := http.Get(signedurl)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Unable to get signed URL")
			return "", err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Unable to get read resp body")
			return "", err
		}
		return string(body), nil
	}

}

func GenerateAmazonUrl(api AmazonMWSAPI, Action string, ActionPath string, Parameters map[string]string) (finalUrl *url.URL, err error) {
	result, err := url.Parse(api.Host)
	if err != nil {
		return nil, err
	}

	result.Host = api.Host
	result.Scheme = "https"
	result.Path = ActionPath

	values := url.Values{}
	values.Add("Action", Action)

	if api.AuthToken != "" {
		values.Add("MWSAuthToken", api.AuthToken)
	}

	values.Add("AWSAccessKeyId", api.AccessKey)
	values.Add("SellerId", api.SellerId)
	values.Add("SignatureVersion", "2")
	values.Add("SignatureMethod", "HmacSHA256")

	if Action == "SubmitFeed" {
		values.Add("Version", "2009-01-01")

		//feed2 := fmt.Sprintf("sku	price\n1000---0734	9.99")
		//sendbody := []byte(feed2)

		//_ := []byte(sendbody)
		//md5Value := fmt.Sprintf("%x", md5.Sum(data))
		//values.Add("Content-MD5", fmt.Sprintf("%x", md5.Sum(data)))

		//params["ContentMD5Value"] = md5Value

	} else if Action == "RequestReport" {
		values.Add("Version", "2009-01-01")
	} else if Action == "GetReportRequestList" {
		values.Add("Version", "2009-01-01")
	} else if Action == "GetReport" {
		values.Add("Version", "2009-01-01")
	} else {
		values.Add("Version", "2011-10-01")
	}

	for k, v := range Parameters {
		values.Set(k, v)
	}

	params := values.Encode()
	result.RawQuery = params

	return result, nil
}

func SetTimestamp(origUrl *url.URL) (err error) {
	values, err := url.ParseQuery(origUrl.RawQuery)
	if err != nil {
		return err
	}
	values.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))
	origUrl.RawQuery = values.Encode()

	return nil
}

func SignAmazonUrl(origUrl *url.URL, api AmazonMWSAPI, Action string) (signedUrl string, err error) {
	escapeUrl := strings.Replace(origUrl.RawQuery, ",", "%2C", -1)
	escapeUrl = strings.Replace(escapeUrl, ":", "%3A", -1)

	params := strings.Split(escapeUrl, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	toSign := fmt.Sprintf("GET\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)
	if Action == "SubmitFeed" {
		toSign = fmt.Sprintf("POST\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)
	}
	if Action == "RequestReport" {
		toSign = fmt.Sprintf("POST\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)
	}
	if Action == "GetReportRequestList" {
		toSign = fmt.Sprintf("POST\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)
	}
	if Action == "GetReport" {
		toSign = fmt.Sprintf("POST\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)
	}

	//toSign := fmt.Sprintf("GET\n%s\n%s\n%s", origUrl.Host, origUrl.Path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(api.SecretKey))
	_, err = hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	newParams := fmt.Sprintf("%s&Signature=%s", sortedParams, hash)

	origUrl.RawQuery = newParams

	return origUrl.String(), nil
}

/*
GetLowestOfferListingsForASIN takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetLowestOfferListingsForASIN(items []string) (string, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceIdDE)

	return api.genSignAndFetch("GetLowestOfferListingsForASIN", "/Products/2011-10-01", params)
}

/*
GetLowestOfferListingsForSKU takes a list of SKUss and returns the result.
*/
func (api AmazonMWSAPI) GetLowestOfferListingsForSKU(items []string) (string, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("SellerSKUList.SellerSKU.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceIdDE)
	params["ExcludeMe"] = "true"

	return api.genSignAndFetch("GetLowestOfferListingsForSKU", "/Products/2011-10-01", params)
}

/*
GetCompetitivePricingForAsin takes a list of ASINs and returns the result.
*/
func (api AmazonMWSAPI) GetCompetitivePricingForASIN(items []string) (string, error) {
	params := make(map[string]string)

	for k, v := range items {
		key := fmt.Sprintf("ASINList.ASIN.%d", (k + 1))
		params[key] = string(v)
	}

	params["MarketplaceId"] = string(api.MarketplaceIdDE)

	return api.genSignAndFetch("GetCompetitivePricingForASIN", "/Products/2011-10-01", params)
}

/*
GetLowestPricedOffersForSKU takes a list of SKUss and returns the result.
*/
func (api AmazonMWSAPI) GetLowestPricedOffersForSKU(sku string) (string, error) {
	params := make(map[string]string)

	//for k, v := range items {
	//	key := fmt.Sprintf("SellerSKU.%d", (k + 1))
	//	params[key] = string(v)
	//}

	params["SellerSKU"] = sku
	params["MarketplaceId"] = string(api.MarketplaceIdDE)
	params["ItemCondition"] = "used"

	return api.genSignAndFetch("GetLowestPricedOffersForSKU", "/Products/2011-10-01", params)
}

func (api AmazonMWSAPI) GetMatchingProductForId(idType string, idList []string) (string, error) {
	params := make(map[string]string)

	for k, v := range idList {
		key := fmt.Sprintf("IdList.Id.%d", (k + 1))
		params[key] = string(v)
	}

	params["IdType"] = idType
	params["MarketplaceId"] = string(api.MarketplaceIdDE)

	return api.genSignAndFetch("GetMatchingProductForId", "/Products/2011-10-01", params)
}

func (api AmazonMWSAPI) SubmitFeed(feedType, feed string) (string, error) {
	params := make(map[string]string)

	params["FeedType"] = feedType
	//params["FeedContent"] = feed
	//params["Attachment"] = feed
	params["MarketplaceIdList.Id.1"] = string(api.MarketplaceIdDE)

	//generate the md5hash

	//data := []byte(feed)
	//md5Value := base64.StdEncoding.EncodeToString(fmt.Sprintf("%x", md5.Sum(data))

	//hasher := md5.New()
	//hasher.Write([]byte(feed))
	//md5Value := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	var outputFileName string
	if os.Args[1] == "DE" {
		outputFileName = "/tmp/newPricesDE.txt"
	} else {
		outputFileName = "/tmp/newPrices.txt"
	}

	md5Value, err := util.Hash_file_md5(outputFileName)
	if err == nil {
		fmt.Println(" Hash is ******************")
		fmt.Println(md5Value)
		fmt.Println(" End Hash is ******************")

	}

	params["ContentMD5Value"] = md5Value

	return api.genSignAndFetch("SubmitFeed", "/Feeds/2009-01-01", params)
}

func (api AmazonMWSAPI) RequestReport(reportType string) (string, error) {
	params := make(map[string]string)

	params["MarketplaceIdList.Id.1"] = string(api.MarketplaceIdDE)
	params["ReportType"] = reportType

	fmt.Printf("len params = %v /n", params)

	return api.genSignAndFetch("RequestReport", "/Reports/2009-01-01", params)
}

func (api AmazonMWSAPI) GetReportRequestList(reportRequestId string) (string, error) {
	params := make(map[string]string)

	params["ReportRequestIdList.Id.1"] = reportRequestId

	return api.genSignAndFetch("GetReportRequestList", "/Reports/2009-01-01", params)
}

func (api AmazonMWSAPI) GetReport(generatedReportId string) (string, error) {
	params := make(map[string]string)

	params["ReportId"] = generatedReportId

	return api.genSignAndFetch("GetReport", "/Reports/2009-01-01", params)
}
