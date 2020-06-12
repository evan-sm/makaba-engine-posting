package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func makeClient() (*http.Client, bool) {
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	auth := CurrentUsercode.PasscodeAuth()
	if auth == false {
		log.Println("Failed to authorize passcode. Skip.")
	}
	cookie := &http.Cookie{
		Name:   "passcode_auth",
		Value:  CurrentUsercode.Usercode,
		Path:   "/",
		Domain: "2ch.hk",
	}
	cookies = append(cookies, cookie)
	u, _ := url.Parse(postingUrl)
	jar.SetCookies(u, cookies)
	//log.Println(jar.Cookies(u))
	client := &http.Client{
		Jar: jar,
	}
	return client, auth
}

func main() {
	client, ok := makeClient()
	if ok == false {
		return
	}
	//prepare the reader instances to encode
	valuesBase := map[string]io.Reader{
		"task":   strings.NewReader("post"),
		"board":  strings.NewReader("test"),  // https://2ch.hk/test/
		"thread": strings.NewReader("28394"), // https://2ch.hk/test/res/28394.html
		//"name": strings.NewReader("#s:|ZX#`j"), // Tripcode for attention whore
		//"email": strings.NewReader(""), // R u fucking kidding me?
		//"subject": strings.NewReader(""), // Oldfags never use it
		"comment": strings.NewReader("Пук"), // Post text
	}
	valuesFiles := map[string]io.Reader{
		`files`: mustOpen("barenzi.jpg"),
		//`files2`: mustOpen("3.jpg"), // lets assume its this file
	}
	err := makabaPost(client, postingUrl, valuesBase, valuesFiles)
	if err != nil {
		log.Println(err)
	}
}

func makabaPost(client *http.Client, url string, valuesBase map[string]io.Reader, valuesFiles map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range valuesBase {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	for key, r := range valuesFiles {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	w.Close()

	// Prepare handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Высрать в тред
	res, err := client.Do(req)
	if err != nil {
		log.Println("client.Do(req) error:", err)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ioutil.ReadAll error:", err)
		return
	}
	log.Println(string(body))
	//var result map[string]interface{}
	//json.NewDecoder(res.Body).Decode(&result)
	//log.Println(result)
	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		log.Println(err)
	}
	return r
}
