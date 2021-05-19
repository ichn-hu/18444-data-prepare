package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql包
	"math/rand"
	"strings"
	"sync"
	"time"
)

var BATCH_SIZE = 1000

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().Unix())
}

func genSingle() string {
	accno := RandStringRunes(10)
	currency := "cny"
	cuno := rand.Int31()
	return fmt.Sprintf("('%s', '%s', %d)", accno, currency, cuno)
}

func genBatch() string {
	batch := make([]string, 0, BATCH_SIZE)
	for i := 0; i < BATCH_SIZE; i++ {
		batch = append(batch, genSingle())
	}
	return "insert into account (accno, currency, cuno) values " + strings.Join(batch, ", ")
}

func main() {
	conc := 10
	num := 10000
	wg := sync.WaitGroup{}
	for i := 0; i < conc; i++ {
		wg.Add(1)
		go func(i int) {
			db, err := sql.Open("mysql", "root:@tcp(0.0.0.0:4201)/test")
			if err != nil {
				fmt.Println(err)
			}
			defer db.Close()

			for j := 0; j < num; j++ {
				_, err := db.Exec(genBatch())
				if err != nil {
					fmt.Println(err)
				}
				if j%10 == 0 {
					fmt.Printf("%d: %d\n", i, j)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("vim-go")
}
