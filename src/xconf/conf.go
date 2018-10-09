package xconf

import (
	"fmt"
	"os"

	"github.com/cihub/seelog"
)

//LogLevel seelog Level
type LogLevel int

const (
	//Critical for critical
	Critical LogLevel = iota
	//Error for error,critical
	Error
	//Warn for warn, error,critical
	Warn
	//Info for info,warn, error,critical
	Info
	//Debug for debug,info,warn, error,critical
	Debug
	//Trace for trace,debug,info,warn, error,critical
	Trace
)

func (r LogLevel) String() string {

	switch r {
	case Critical:
		return "critical"
	case Error:
		return "error,critical"

	case Warn:
		return "warn, error,critical"

	case Info:
		return "info,warn, error,critical"

	case Debug:
		return "debug,info,warn, error,critical"

	case Trace:
		return "trace,debug,info,warn, error,critical"
	default:
		return "warn,error,critical"
	}
}

func getSeelogDefaultHourConf(logDir string, logName string, console bool, level LogLevel) string {
	if console {
		return `<seelog minlevel="trace">
    <outputs formatid="main">   	
        <filter levels="` + level.String() + `">           
			<rollingfile type="date" filename="` + logDir + logName + `.txt" datepattern="2006.01.02-15" maxrolls="7" />
        </filter>
		<filter levels="` + level.String() + `"> 
			<console />
		 </filter>
    </outputs>
    <formats>
        <format id="main" format="[%Filename %Line] %Date(2006-01-02T15:04:05.999999999Z07:00) [%LEV] %Msg%n"/>    -->format内容，可以多个共存，只要id不相同。然后上面可以用不同的id来输出不同格式的日志。
    </formats>
</seelog>
`
	}
	return `<seelog minlevel="trace">
    <outputs formatid="main">   	
        <filter levels="` + level.String() + `">           
			<rollingfile type="date" filename="` + logDir + logName + `.txt" datepattern="2006.01.02-15" maxrolls="7" />
        </filter>		
    </outputs>
    <formats>
        <format id="main" format="[%Filename %Line] %Date(2006-01-02T15:04:05.999999999Z07:00) [%LEV] %Msg%n"/>    -->format内容，可以多个共存，只要id不相同。然后上面可以用不同的id来输出不同格式的日志。
    </formats>
</seelog>
`
}

func getSeelogDefaultDayConf(logDir string, logName string, console bool, level LogLevel) string {
	if console {
		return `<seelog minlevel="trace">
    <outputs formatid="main">   	
        <filter levels="` + level.String() + `">           
			<rollingfile type="date" filename="` + logDir + logName + `.txt" datepattern="2006.01.02" maxrolls="7" />
        </filter>
		<filter levels="warn,error,critical"> 
			<console />
		 </filter>
    </outputs>
    <formats>
        <format id="main" format="[%Filename %Line] %Date(2006-01-02T15:04:05.999999999Z07:00) [%LEV] %Msg%n"/>    -->format内容，可以多个共存，只要id不相同。然后上面可以用不同的id来输出不同格式的日志。
    </formats>
</seelog>
`
	}
	return `<seelog minlevel="trace">
    <outputs formatid="main">   	
        <filter levels="` + level.String() + `">           
			<rollingfile type="date" filename="` + logDir + logName + `.txt" datepattern="2006.01.02" maxrolls="7" />
        </filter>		
    </outputs>
    <formats>
        <format id="main" format="[%Filename %Line] %Date(2006-01-02T15:04:05.999999999Z07:00) [%LEV] %Msg%n"/>    -->format内容，可以多个共存，只要id不相同。然后上面可以用不同的id来输出不同格式的日志。
    </formats>
</seelog>
`
}

func getSeelogConfFile(confDir string, logName string, hour bool, console bool) string {
	if hour {
		if console {
			return confDir + logName + "_seelog_console_hour.xml"
		}
		return confDir + logName + "_seelog_file_hour.xml"
	}
	if console {
		return confDir + logName + "_seelog_console_day.xml"
	}
	return confDir + logName + "_seelog_file_day.xml"
}

//InitSeelogContext 初始化seelog 所需要的参数
type InitSeelogContext struct {
	ConfDir      string   //配置文件所在目录
	LogDir       string   //日志文件目录
	logName      string   //文件名称，用于配置文件和日志输出文件
	Hour         bool     //是否按小时输出日志文件，false 为按天，ture按小时
	Console      bool     //是否输出到控制台, titan 启动 flutend需要设置为true
	FileLevel    LogLevel //日志输出到文件时的级别
	ConsoleLevel LogLevel //日志输出到控制台的级别
}

//InitSeelog 初始化进程的seelog配置
//confDir 配置文件所在目录， eg  /ssd/repro_go/conf/
//logDir  日志文件保存的目录 , eg /ssd/repro_go/log/
//logName  日志文件的名称(建议与docker 容器名称相同)， eg  repro_go1.txt  , 此时的文件全路径为 logDir+logName eg: /ssd/repro_go/log/repro_go1.txt
//hour   是否按消息滚动日志文件， true，每个小时一个日志文件，否则每天一个日志文件
//console enable output log to STDOUT
func InitSeelog(confDir string, logDir string, logName string, hour bool, console bool, level LogLevel) {
	//seelog配置文件路径，从reprosrv的配置文件读取位置
	logConfigFile := getSeelogConfFile(confDir, logName, hour, console)
	fmt.Println("seelog conf:", logConfigFile)
	//加载seelog配置，配置文件的内容一定要正确。
	logger, err := seelog.LoggerFromConfigAsFile(logConfigFile)

	if err != nil {
		//如果初始化失败，则创建默认的seelog配置文件
		fmt.Printf("load seelog from file %s error %s; retry", logConfigFile, err)
		file, err := os.OpenFile(logConfigFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		if hour {
			file.WriteString(getSeelogDefaultHourConf(logDir, logName, console, level))
		} else {
			file.WriteString(getSeelogDefaultDayConf(logDir, logName, console, level))
		}

		file.Close()

		logger, err = seelog.LoggerFromConfigAsFile(logConfigFile)
		if err != nil {
			fmt.Printf("load seelog from file %s error %s; exit", logConfigFile, err)
			return
		}
	}
	seelog.ReplaceLogger(logger)

}
