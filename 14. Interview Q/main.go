package main

import (
	"fmt"
	"sync"
)

// 3 table
func PerformDataOperations(tablesName []string) bool {
	// To synchronize the data without blocking the goroutine:- array of fixed size of tableNames, and
	//send index in arg in every go routine and then insert data in the array based on that goroutine
	//We can also do it by using map
	finalAns := ""
	var wg sync.WaitGroup
	wg.Add(len(tablesName))
	ch := make(chan string, len(tablesName))
	for _, tableName := range tablesName {
		go ReadDataFromTable(tableName, ch, &wg)
	}
	wg.Wait()
	close(ch)

	for msg := range ch {
		finalAns += msg
	}

	ch2 := make(chan bool, len(tablesName))
	var wg2 sync.WaitGroup
	wg2.Add(len(tablesName))
	for _, tableName := range tablesName {
		go WriteDataToTable(tableName, ch2, finalAns, &wg2)
	}
	wg2.Wait()
	close(ch2)

	for msg := range ch2 {
		if !msg {
			return false
		}
	}
	return true

}

func ReadDataFromTable(tableName string, ch chan string, wg *sync.WaitGroup) {
	//Lock
	ch <- "data from:" + tableName
	wg.Done()
}

func WriteDataToTable(tableName string, ch chan bool, data string, wg *sync.WaitGroup) error {
	fmt.Println("Writing data to table", tableName, data)
	ch <- true
	wg.Done()
	return nil
}

func main() {
	tableName := []string{"table1", "table2", "table3"}
	fmt.Print(PerformDataOperations(tableName))
}

/*
We have to implement a function "PerformDataOperations". Please find below the specifics:
Signature:
func PerformDataOperations(tablesName []string) bool
3 tables
data1
data2
data2

data1data2data3 -> t1, t2, t3

Specification/Steps:
Step1:
It should perform read operations on the input tables list.
The data read from the tables is of type string
All the read operations are independant of each other

Step2:
After all the data from the tables is read, we will process the data and prepare "result".
This "result" will be of type string.
E.g: For 2 tables , table1 and table2. Lets say we get records from these tables as "record1", "record2" respectively
Than the result will be:
result ="record1record2"

Step3: Once the data is processed and we have the "result", we will write this data back to input tables list.
Each write operation is independant of each other.

Step4: Once the write operations are complete, the function will return
    true: If all the write operations succeeded
    false: If any write operation failed


Note: For the read and write operations, we don't need to write actual db query.
We can write the stubs for these.

Sample stubs:
func ReadDataFromTable(tableName string) string {
    return "data from:"+tableName
}

func WriteDataToTable(tableName string, data string) error {
    return nil
}

Note: These are only samples for reference and are not complete.
You will have to write the function signature as per your logic
*/
