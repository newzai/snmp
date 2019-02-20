package xwarning

import "github.com/go-xorm/xorm"

type operatorType int

const (
	cGetCount operatorType = iota
	cAdd
	cDel
)

type message struct {
	tid          int
	opType       operatorType
	responseChan chan int
}

//WarningStats   设备告警个数缓存
type WarningStats struct {
	engine *xorm.Engine
	counts map[int]int // key is terminal id (tid),value is warning counts
	opChan chan *message
}

func newWarningStats(engine *xorm.Engine) *WarningStats {
	return &WarningStats{
		engine: engine,
		counts: make(map[int]int),
		opChan: make(chan *message, 100),
	}
}

func (r *WarningStats) exec() {

	for msg := range r.opChan {

		switch msg.opType {
		case cGetCount:
			value, ok := r.counts[msg.tid]
			if ok {
				msg.responseChan <- value
			} else {
				warngs, _ := GetWarningsByTID(msg.tid, r.engine)
				r.counts[msg.tid] = len(warngs)
				msg.responseChan <- len(warngs)
			}
		case cAdd:
			value, ok := r.counts[msg.tid]
			if ok {
				value++
				r.counts[msg.tid] = value
			}
		case cDel:
			value, ok := r.counts[msg.tid]
			if ok && value > 0 {
				value--
				r.counts[msg.tid] = value
			}
		}
	}
}

//GetCounts return tid warning counts
func (r *WarningStats) GetCounts(tid int) int {

	msg := &message{
		tid:          tid,
		opType:       cGetCount,
		responseChan: make(chan int, 1),
	}

	defer close(msg.responseChan)
	r.opChan <- msg

	value := <-msg.responseChan
	return value
}

//Add add warnings for tid
func (r *WarningStats) Add(tid int) {
	msg := &message{
		tid:    tid,
		opType: cAdd,
	}
	r.opChan <- msg
}

//Del delelete count for tid
func (r *WarningStats) Del(tid int) {
	msg := &message{
		tid:    tid,
		opType: cDel,
	}
	r.opChan <- msg
}

//Stats defaulut
var Stats *WarningStats
