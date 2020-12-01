package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func dao(id int) (string, error) {
	return "sdadadas", sql.ErrNoRows
}

func service(id int) (string, error) {
	content, err := dao(id)
	isImportant := func(content int) bool {
		if id < 300 {
			return true
		}
		return false
	}
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// 这里根据业务的需要进行操作，如果是像很重要的业务
			// 就要将这个错误error向上抛
			// 如果这个业务可以支持降级等，就不需要抛，在本层处理好就可以

			// 内容不重要
			if isImportant(id) {
				return "Default Content", nil
			}
		}
		return content, err
	}
	// process the content, return the content when no error
	return content, nil
}

func main() {
	id := 123456
	content, err := service(id)
	if err != nil {
		log.Printf("error: \n%+v\n", err)
		fmt.Println("500")
		return
	}
	fmt.Println("200", content)
}
