package allocateid

import (
	"testing"
)

func Test_AllocateID(t *testing.T) {

	for i := 0; i < 10; i++ {
		AllocateID()
		AllocateID()
	}

}
