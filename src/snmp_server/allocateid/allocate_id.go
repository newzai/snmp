package allocateid

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

var ids []int
var mutex sync.Mutex

//AllocateID allocate id
func AllocateID() int {

	mutex.Lock()
	defer mutex.Unlock()

	if len(ids) > 0 {
		ret := ids[0]
		ids = ids[1:]
		return ret
	}

	data, err := ioutil.ReadFile("./allocate_id")
	if err != nil {
		for i := 0; i < 10; i++ {
			ids = append(ids, i+1)
		}
	} else {

		start, err := strconv.Atoi(string(data))

		if err != nil {
			panic(err)
		}

		for i := start; i < 10+start; i++ {
			ids = append(ids, i+1)
		}
	}

	end := fmt.Sprintf("%d", ids[9])
	ioutil.WriteFile("./allocate_id", []byte(end), os.ModePerm)

	ret := ids[0]
	ids = ids[1:]
	return ret
}
