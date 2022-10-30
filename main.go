package main

import (
	"io"
	"log"
	"os"

	"github.com/NVIDIA/go-tfdata/tfdata/core"
)

func prepareExamples(cnt int) []*core.TFExample {
	result := make([]*core.TFExample, 0, cnt)
	for i := 0; i < cnt; i++ {
		ex := core.NewTFExample()
		ex.AddIntList("int-list", []int{0, 1, 2, 3, 4, 5})
		ex.AddFloat("float", 0.42)
		ex.AddBytes("bytes", []byte("bytesstring"))
		result = append(result, ex)
	}

	return result
}

func writeExamples(w io.Writer, examples []*core.TFExample) error {
	tfWriter := core.NewTFRecordWriter(w)
	for _, example := range examples {
		_, err := tfWriter.WriteExample(example)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	path := "testtfrecordwriter.record"
	sinkFd, err := os.Create(path)

	if err != nil {
		log.Fatalln(err)
	}

	recordWriter := core.NewTFRecordWriter(sinkFd)

	results := prepareExamples(1)
	log.Println(results)

	writeExamples(recordWriter, results)

}
