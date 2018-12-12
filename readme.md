<!-- TOC -->

- [1. build](#1-build)
- [2. common](#2-common)
- [3. login](#3-login)
- [4. logout](#4-logout)
- [5. getitem](#5-getitem)
- [6. SNMP 配置](#6-snmp-配置)
    - [6.1. 获取配置文件](#61-获取配置文件)
- [7. SNMP 批量配置](#7-snmp-批量配置)
    - [7.1. 批量FTP升级](#71-批量ftp升级)
    - [7.2. 批量重启](#72-批量重启)
    - [7.3. 批量恢复出厂设置](#73-批量恢复出厂设置)
    - [7.4. 批量备份配置文件](#74-批量备份配置文件)
    - [7.5. 批量配置](#75-批量配置)
- [8. modify_password  修改自己的用户名和密码](#8-modify_password--修改自己的用户名和密码)
- [9. getusers  -- for admin](#9-getusers-----for-admin)
- [10. create_user -- for admin](#10-create_user----for-admin)
- [11. delete_user -- for admin](#11-delete_user----for-admin)
- [12. modify_user -- for admin](#12-modify_user----for-admin)
- [13. 固件管理  -- for admin](#13-固件管理-----for-admin)
    - [13.1. 固件查询](#131-固件查询)
    - [13.2. 固件上传](#132-固件上传)
- [14. FTP 固件升级](#14-ftp-固件升级)
    - [14.1. ftp 升级](#141-ftp-升级)
    - [14.2. ftp 升级状态查询](#142-ftp-升级状态查询)
- [15. linux 命令](#15-linux-命令)
    - [15.1. 获取命令列接口](#151-获取命令列接口)
    - [15.2. 执行命令](#152-执行命令)
- [16. 内嵌web](#16-内嵌web)
    - [16.1. dist.zip](#161-distzip)
    - [16.2. build 内嵌资源 asset.go 文件](#162-build-内嵌资源-assetgo-文件)
    - [16.3. go-bindata 安装 (执行一次即可)](#163-go-bindata-安装-执行一次即可)

<!-- /TOC -->

# 1. build 

```bash
export GOPATH=`pwd`
cd src/snmp_server
make

./snmp_server -version
```

# 2. common
- userinfo 
```json
{
    "userid":int
    "usertype":int  -- 1 管理员， 2 普通用户
    "username": ""
    "password": ""  -- md5之后的值
    zoneinfo:
    {
        "zoneid": int   -- top(root) zone
        "zonename": ""  -- top(root) zone name
        "zonepath": ""  
    }
}
```
- iteminfo
```json
{
    
    "itemid":int
    "parent":int
    "itemname":""
    "itempath":""
    "itemtype":int 
    "status": int  -- for itemtype == 2
    "dev_type":"" -- 设备类型 for itemtype == 2
            
}
```
# 3. login 
- POST
- URL v1/login
- Body
```json
{
    "username":"",
    "password":md5("")
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok
    "message": "OK"     -- or other error message
    "data":             -- if result == 0
    {
        "token":""      --登录成功后分配的token值，后续请求都需要
        "userinfo": userinfo
    }
}
```
- Test  

curl -X POST -H 'content-type: application/json' -d '{"username":"admin", "password":"123456"}' http://118.126.91.183:9192/v1/login


# 4. logout 
- POST
- URL v1/logout
- Body
```json
{
    "token":""     -- login 返回的token
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok
    "message": "OK"     -- or other error message
}
```

# 5. getitem  

curl -X POST -H 'content-type: application/json' -d '{"token":"111-222-333-4444", "data":{"itemid":1, "itemtype":1}}' http://127.0.0.1:9192/v1/getitem

- POST
- URL v1/getitem
- Body
```json
{
    "token":""
    "data":
    {
        "itemid": int  -- zoneid or devid or userid
        "itemtype": int -- 1 zone, 2 dev , 3 user
    }
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
    "data":
    {
        "items":
        [
            iteminfo_1,
            iteminfo_2,
            ...
            iteminfo_i
        ]
    }
}
```
- Test 

curl -X POST -H 'content-type: application/json' -d '{"token":"111-222-333-4444", "data":{"itemid":1, "itemtype":1}}' http://118.126.91.183:9192/v1/getitem

    
# 6. SNMP 配置

- POST
- URL v1/snmp
- Body
```json
{
    "token":""
    "data":
    {
        "snmp_type": ""  -- value is "get" or "set"
        "itemid":int     --  要操作的设备id 
        "index" : int    --- 一般填写0， 操作snmp table 填写索引值
        "oids":
        {
            一下成员，动态，根据具体页面设置即可, 对于get int填写0， string 填写 "" 
            "oid1": int or str
            "oid2" :int or str
            ....
            "oidi": int or str
        }
    }
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
    "data":             -- only for snmp_type is get 
    {
        "oids":
        {
            "oid1": int or str
            "oid2" :int or str
            ....
            "oidi": int or str
        }
    }
}
```
- Test

 curl -X POST -H 'content-type: application/json' -d '{"token":"111-222-333-4444", "data":{"snmp_type":"get", "itemid":4,"index":1,"oids":{"software_version":"","wan_link_status":0,"lan_link_status":0}}}'  http://118.126.91.183:9192/v1/snmp
 
## 6.1. 获取配置文件 
- SNMP 配置更新之前，可以获取配置文件列表
- POST 
- URL v1/get_config_file
- Body 
```json
{
    "token":""
    "data":
    {
        "itemid": 5
    }
}
```
- Response 
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
    "data":
    {
        "files":{
            file1,
            file2,
            ...,
            fileN
        }
    }
}
```
- 把获取到的files的 任意一个文件名称放到 oids的字段 usl_ftp_restore_cfg_file_name 即可


# 7. SNMP 批量配置
- 说明，只能批量配置，归属自己的设备
- POST
- URL v1/snmp_batch
- Body
```json
{
    "token":""
    "data":
    {
        "dev_type": "" --要批量配置的设备类型
        "itempath": "" -- 指定 zone path ，在该path下面的所有指定设备会被批量处理
        "itemids": [id1,id2...], --优先，len(itemids) >0 忽略itempath
        "oids":
        {
            "oid1": int or str
            "oid2" :int or str
            ....
            "oidi": int or str
        }
    }
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
    "data":
    {
        "ok_items":   --所谓的成功是直snmp命令发送成功， 如果用于固件 ftp升级，不代表固件ftp升级成功，只能说是把升级命令发送到了设备
        [
            {
                "itemid":int,
                "itemname":""
            }
        ]
        "error_items":
        [
           {
               "itemid":int
               "itemname":""
               "error":""
           }
        ],

    }
}
```
## 7.1. 批量FTP升级
包含ftp升级相关参数 

```json
{
    "oids":
    {
        "usl_ftp_server_ip":"118.126.91.183",
        "usl_ftp_server_port":21,
        "usl_ftp_user_name":"uftp",
        "usl_ftp_user_passwd":"123456789",
        "usl_ftp_file_size":"380605",
        "usl_ftp_soft_file_name":"klph-40b19.tar"
    }
}
```  
            
## 7.2. 批量重启
```json
{
    "oids":
    {
        "usl_reboot_device":"reboot"
    }
}
```

## 7.3. 批量恢复出厂设置
```json
{
    "oids":
    {
        "usl_set_default":"default"
    }
}
```
## 7.4. 批量备份配置文件
```json
{
    "oids":
    {
        "usl_ftp_server_ip":"118.126.91.183",
        "usl_ftp_server_port":21,
        "usl_ftp_user_name":"uftp",
        "usl_ftp_user_passwd":"123456789",
        "usl_ftp_save_cfg_file_name":"default"
    }
}
```  

usl_ftp_save_cfg_file_name 名称随意给即可，服务器会统一文件名称为 设备的ntid_YYmmDD_HHMMSS.cfg
## 7.5. 批量配置

以下列出的字段支持批量配置，可以出现在 请求消息 的oids参数中；  

|名称|jsonKey|默认值|说明| 
|----|:----:|:----|:---:|
| 语言模式 |	k518_language_mode |               缺省：0 |   //中文 |  
| http模式 |	k518_http_mode   |                 缺省：0 |  //http模式 |  
|http端口 |	k518_http_port  |                  缺省：80  |  |
|telnet端口 |	k518_telnet_port   |               缺省：23  |   |
|sip本地端口 |	k518_sip_local_port  |             缺省：5060   |   |
|二次拨号模式|	k518_dtmf_relay_mode |             缺省：0  |//2833模式  |
|语音算法列表|	k518_codec_type_list   |           缺省："3 1 0 4 2"  ||
|热线使能开关|k518_hotline_enable   |            缺省：0   |//关闭   |
|热线号码	|k518_hotline_number   |            缺省："690" |  |
|热线账号|	k518_hotline_account|              缺省：0     | |
|麦克风音量	| k518_microphone_volume |           缺省：6    | | 
|扬声器音量 |	k518_speaker_volume         |      缺省：4   | |  
|挂机等待时间	|k518_hookon_wait_time  |           缺省：10  |   
|铃声类型	|k518_ring_style     |              缺省：0    |
|振铃音量	|k518_ring_volume     |             缺省：7     |
|语言模式	|k519_language_mode    |            缺省：0      |                  //中文
|http模式	|k519_http_mode       |             缺省：0     |                   //http模式
|http端口	|k519_http_port       |             缺省：80   |
|telnet端口	|k519_telnet_port        |          缺省：23   |
|sip本地端口	|k519_sip_local_port       |        缺省：5060   |
|二次拨号模式	|k519_dtmf_relay_mode    |          缺省：0        |                //2833模式
|语音算法列表	|k519_codec_type_list     |         缺省："3 1 0 4 2"|
|热线使能开关	|k519_hotline_enable      |         缺省：0       |                 //关闭
|热线号码	|k519_hotline_number          |     缺省："690" |
|热线账号	|k519_hotline_account          |    缺省：0   |
|视频模式	|k519_video_mode	    	  |        缺省：0    |                    //720p
|视频负载类型	|k519_video_paytype	       |   缺省：107       |               
|麦克风音量	|k519_microphone_volume         |   缺省：6   |
|扬声器音量	|k519_speaker_volume            |   缺省：4   |
|挂机等待时间	|k519_hookon_wait_time       |      缺省：10  | 
|铃声类型	|k519_ring_style                |   缺省：0  |
|振铃音量	|k519_ring_volume                |  缺省：7   |

# 8. modify_password  修改自己的用户名和密码
- POST
- URL v1/modify_password
- Body
```json
{
    "token":""
    "data":
    {
        "userid": int -- user for modify
        "old_password":""  -- md5
        "new_password": "" -- md5
    }
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
}
```

# 9. getusers  -- for admin 
- POST
- URL v1/getusers
- Body
```json
{
    "token":""
    "data":
    {
        "zoneid": int 
    }
}
```



- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
    "data":
    {
        "users":[
        userinfo_1,
        userinfo_2,
        ...
        userinfo_i
        ]
    }
}
```
# 10. create_user -- for admin
- POST 
- URL  v1/create_user
- Body
```json
{
    "token":"",
    "data":
    {
        "username":""
        "password": "" -- md5
        "usertype": int
        "zoneid":int -- parent zone id 
    }
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
}
```
# 11. delete_user -- for admin
- POST
- URL v1/delete_user
- Body
```json
{
    "token":"",
    "data":
    {
        "userid":int -- 要删除的用户id
    }
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
}
```

# 12. modify_user -- for admin 
- POST
- URL v1/modify_user
- Body
```json
{
    "token":"",
    "data":
    {
        "userid": int  -- for modify user id 
        "usertype": int  --如果要修改用户类型，包含该字段
        "password": "" -- md5  如果要修改 密码，包含该字段
        "zoneid": int  -- 如果要修改管理的区域，包含该字段
    }
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
}
```

# 13. 固件管理  -- for admin
## 13.1. 固件查询

- POST
- URL  v1/get_all_hardware
- Body
```json
{
    "token":""
    "data":{
        "dev_type":"KN518"| "KN519"
    }
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
    "data":
    {
        "files":
        [
            filename1,
            filename2,
            ....
            filenameN
        ]
    }
}
```

## 13.2. 固件上传

通过  From 表单上传  





# 14. FTP 固件升级
## 14.1. ftp 升级

- 使用 snmp set命令，参考 snmp接口；需要包含以下字段
    usl_ftp_server_ip  
    usl_ftp_server_port  
    usl_ftp_user_name  
    usl_ftp_user_passwd  
    usl_ftp_file_size  
    usl_ftp_soft_file_name  
- URL  v1/snmp

```json
{
	"token":"111-222-333-4444",
	"data":
	{
		"snmp_type":"set",
		"itemid":5,
		"index":0,
		"oids":
		{
			"usl_ftp_server_ip":"118.126.91.183",
			"usl_ftp_server_port":21,
			"usl_ftp_user_name":"uftp",
			"usl_ftp_user_passwd":"123456789",
			"usl_ftp_file_size":"380605",
			"usl_ftp_soft_file_name":"klph-40b19.tar"
		}
	}
}
```
## 14.2. ftp 升级状态查询

- POST 
- URL /v1/ftp_upgrade_status
- Body 
```json
{
    "token":
    "data:"
    {
        itemid: int  -- 查询单个设备
        itempath: string --查询批量ftp升级
        dev_type: string --查询批量ftp升级
    }
}
```
- Response
```json
{
    "result": 0         -- or other, 0 is ok, if ok has data
    "message": "OK"     -- or other error message
    "data":
    {
       "upgrades": [ftp_upgrade_info...]
    }
}
```

- ftp_upgrade_info
```json
{
    itemid: int 
    itemname: string
    itempath: string
    dev_type: string
    result :
}
```


# 15. linux 命令

- 为了安全，只能执行指定的系统命令。相关系统命令放在与snmp_server 相同的目录 command.txt文件中。
- command.txt 文件每行一个命令，每一行数据保护三个字段， index:name:command
- index 必须唯一
- name 为名称，用于web页面展示
- command具体执行的内容，对web 和用户透明。

## 15.1. 获取命令列接口

- POST
- URL : /v1/get_commands
- Body 
```json
{
	"token":"8008fd2f-3ccd-4582-8756-ae3d13ea7f77"
}
```

- Response
```json
{
    "data": {
        "commands": {
            "1": "ps",
            "2": "ls",
            "3": "ping",
            "4": "test",
            "5": "test"
        },
        "hash": "6688f785308916906ccce2095814528a"
    },
    "message": "OK",
    "result": 0
}
```

- 把Response结果的 commands 展示在web界面上，属于所有支持的命令,
- KEY 为 index
- Value 为 name
- get_commands 的列表，后台修改后，会导致变化，表现的结果就是hash值不一样。

## 15.2. 执行命令 
- POST
- URL : /v1/run_command
- Body
```json
{
	"token":"952f2cdb-b308-4018-be0c-a7ea196f1f7f",
	"data":{
		"command": 3,
        "params": "",
		"hash":"6688f785308916906ccce2095814528a"
	}
}
```
- Response 
```json
{
    "data": {
        "output": "PING 127.0.0.1 (127.0.0.1) 56(84) bytes of data.\n64 bytes from 127.0.0.1: icmp_seq=1 ttl=64 time=0.046 ms\n64 bytes from 127.0.0.1: icmp_seq=2 ttl=64 time=0.031 ms\n64 bytes from 127.0.0.1: icmp_seq=3 ttl=64 time=0.033 ms\n64 bytes from 127.0.0.1: icmp_seq=4 ttl=64 time=0.036 ms\n\n--- 127.0.0.1 ping statistics ---\n4 packets transmitted, 4 received, 0% packet loss, time 2997ms\nrtt min/avg/max/mdev = 0.031/0.036/0.046/0.008 ms\n"
    },
    "message": "OK",
    "result": 0
}
```

- command 的值为 3， 对比 get_commands 的返回结果得知执行ping命令
- hash get_commands 返回的hash，需要重新带回去，服务器会校验，如果hash值不匹配返回错误，需要重新获取get_commands,因此后端的command.txt 文件被修改了。
- params 可以传递命令需要的相关参数  
- Response hash not match
```json
{
    "message": "hash not match, refresh commands list",
    "result": 2
}
```


# 16. 内嵌web

## 16.1. dist.zip 

- dist.zip 文件放到 src/snmp_server目录
- unzip dist.zip 解压到 src/snmp_server目录， 出现 src/snmp_server/dist 目录 

## 16.2. build 内嵌资源 asset.go 文件

```bash
cd src/snmp_server 
go-bindata -o asset/asset.go -pkg=asset dist/...
```

## 16.3. go-bindata 安装 (执行一次即可)

```bash
export GOPATH=`pwd` #在snmp 目录设置
go get -u github.com/jteeuwen/go-bindata/...
```

- 安装完成后，把 `pwd`/bin 目录加入 PATH目录或者把 `pwd`/bin/go-bindata 文件复制到 /usr/local/bin目录下面

