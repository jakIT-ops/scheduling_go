package Jobs

import (
	"fmt"
	"log"
)

type Jobs struct{}

func (jobs *Jobs) Hello() {
	log.Println("Start Job: Hello")
	fmt.Println("Hello:")
	fmt.Println("10 second tutam ajillana")
	log.Println("End Job: Hello")
}

func (jobs *Jobs) Hello1() {
	log.Println("Start Job: Hello 1")
	fmt.Println("Hello: JAWHAA")
	fmt.Println("5 second tutam ajillana")
	log.Println("End Job: Hello 1")
}

func (jobs *Jobs) Hello2() {
	log.Println("Start Job: Hello 2")
	fmt.Println("Hello: JAKIT")
	fmt.Println("3 second tutam ajillana")
	log.Println("End Job: Hello 2")
}

func (jobs *Jobs) hello3() {
	log.Println("Start Job: Hello 3")
	fmt.Println("Hello: BAYRJAWKHLAN")
	fmt.Println("1 second tutam ajillana")
	log.Println("End Job: Hello 3")
}
