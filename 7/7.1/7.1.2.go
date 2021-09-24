// 7.1.2 고루틴 기다리기
// 따로 노는 고루틴을 제어하기 위해 싱크 라이브러리 제공된다.

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

var urls = []string{
	"http://image.com/img01.jpg",
	"http://image.com/img02.jpg",
	"http://image.com/img03.jpg",
}

func main() {
	for _, url := range urls {
		go func(url string) {
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}

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
		return "", err
	}
	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}
	f, err := os.Create(filename)
	if err != nil {
		return "", err
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
	// 경로에서 가장 마지막 부분을 반환한다. ex) img01.jpg
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

// !! 여기서 문제가 발생, 위 코드는 파일을 다운로들 하는 것과 압축하는 것 중 어느 것을 먼저 수행해도 좋은 동시성 있는 작업이다.
// 그러나 실제로는, 다운로드가 되지 않으면 파일을 압축할 수 없다.
// 이 때, 사용가능 한 것이 " sync.WaitGroup " 이다.

// main의 for 반복문을 이렇게 쓰면 된다.
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(len(urls)) // 몇 개의 고루틴이 생길지 모를 때는 wg.Add(1)을 for문 안에 해준다.
// 	for _, url := range urls {
// 		go func(url string) {
// 			defer wg.Done()
// 			if _, err := download(url); err != nil {
// 				log.Fatal(err)
// 			}
// 		}(url)
// }
// wg.Wait()

// 	filenames, err := filepath.Glob("*.jpg")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = writeZip("images.zip", filenames)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// **** 공유 메모리와 병렬 최소값 찾기 ****
// 고루틴들이 파일시스템에 데이터를 저장하면 기다렸다가 파일 시스템에서 파일 목록을 찾은 다음 사용했다.
// 고루틴은 메모리도 서로 공유하기 때문에, 파일시스템 사용하지 않고 포인터로 값을 바꿔줄 수 있다.

// 일단 병렬화 하지 않고 최소값 찾기
func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}

func ExampleMin() {
	fmt.Println(Min([]int{
		83, 46, 49, 23, 44,
	}))
}

// 병렬화 시킨 버전
func ParallelMin(a []int, n int) int {
	if len(a) < n {
		return Min(a)
	}
	mins := make([]int, n)
	size := (len(a) + n - 1) / n
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i] = Min(a[begin:end])
		}(i)
	}
	wg.Done()
	return Min(mins)
}

func ExampleParallelMin() {
	fmt.Println(ParallelMin([]int{
		83, 46, 49, 23, 44,
	}, 4))
}
