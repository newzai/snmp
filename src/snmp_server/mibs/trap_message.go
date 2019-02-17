package mibs

//IWarning warning interface
type IWarning interface {
	WarningType() string
	WarningStatus() int
	IsClear() bool
	GetNTID() string
	GetDemo() string
}
