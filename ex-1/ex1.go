package ex_1
import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func startTimer(cnt chan bool) {
	ticker := time.NewTicker(10 * time.Second)
	select {
	case t := <-ticker.C:
		fmt.Println("TIME FINISHED", t)
		ticker.Stop()
		cnt <- false
	}
}

func StartQuiz() {
	csvfile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("can not open file", csvfile)
	}
	r := csv.NewReader(csvfile)
	var correct, incorrect int = 0, 0
	cont := make(chan bool)
	go startTimer(cont)
	go func() {
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("what is   %s, Answer: ", record[0])

			ans, _ := reader.ReadString('\n')
			ans = strings.Replace(ans, "\n", "", -1)
			if ans == record[1] {
				correct++
				fmt.Println("True answer")
			} else {
				incorrect++
				fmt.Println("False answer")
			}

		}
	}()
	<-cont
	fmt.Println("----------------  RESULT ------------------")
	fmt.Printf("Correct answer: %d\n", correct)
	fmt.Printf("InCorrect answer: %d\n", incorrect)
}