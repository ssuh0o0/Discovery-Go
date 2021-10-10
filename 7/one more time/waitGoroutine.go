package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

var urls = []string{
	"https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.pizzaalvolo.co.kr%2Fstatic%2Fassets%2Fimg%2Fimg_alvoloSns_tit.png&imgrefurl=https%3A%2F%2Fwww.pizzaalvolo.co.kr%2F&tbnid=VM2KUSMPQMS8VM&vet=12ahUKEwj_ytD1pL_zAhUYAaYKHQe-CuwQMygBegUIARCZAQ..i&docid=XmMuTQnYvAbE8M&w=720&h=600&q=%ED%94%BC%EC%9E%90%20%EC%95%8C%EB%B3%BC%EB%A1%9C&ved=2ahUKEwj_ytD1pL_zAhUYAaYKHQe-CuwQMygBegUIARCZAQ",
	"https://www.google.com/imgres?imgurl=http%3A%2F%2Fimage.edaily.co.kr%2Fimages%2Fphoto%2Ffiles%2FNP%2FS%2F2021%2F06%2FPS21062200172.jpg&imgrefurl=https%3A%2F%2Fwww.edaily.co.kr%2Fnews%2Fread%3FnewsId%3D01482566629084344%26mediaCodeNo%3D257&tbnid=1Ps1XigClEFc5M&vet=12ahUKEwj_ytD1pL_zAhUYAaYKHQe-CuwQMygEegUIARCfAQ..i&docid=6nNvm9RxxPOUaM&w=441&h=549&q=%ED%94%BC%EC%9E%90%20%EC%95%8C%EB%B3%BC%EB%A1%9C&ved=2ahUKEwj_ytD1pL_zAhUYAaYKHQe-CuwQMygEegUIARCfAQ",
}

func main() {
	var wg sync.WaitGroup
	wg.Add(len(urls)) // 몇 개의 고루틴이 생길지 모를 때는 wg.Add(1)을 for문 안에 해준다.
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait()

	filenames, err := filepath.Glob("*.jpg")
	if err != nil {
		log.Fatal(err)
	}

	err = writeZip("images.zip", filenames)

	if err != nil {
		log.Fatal(err)
	}
}

func download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return filename, err
}

func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return filepath.Base(url.Path), nil

}

func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}
