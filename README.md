# sofa

`sofa`主要用于执行 `shell` 以及`命令行应用` 并收集其输出结果（Console Stdout）并存储到`Redis`,便于后续处理（Logstash）, 
弥补`filebeat`只能收集`log`文件，无法收集控制台输出到问题，也可以通过`Redis`任务队列来接受命令执行并处理结果

主要有以下几个命令：
- `sofa`: 
    通过linux管道接受`shell` 以及`命令行应用` 的输出结果，通过 `unix socket` 发送给`sofa server`来处理并存储到`Redis`

    e.g:
    
        docker ps | sofa -p docker -c ps -t 001
    
- `sofa exec`:
    直接执行一条命令并获取输出结果，通过 `unix socket` 发送给`sofa server`来处理并存储到`Redis`
    
     e.g:
        
        sofa exec "docker ps" -p docker -c ps -t 001
        
        
- `sofa server`:
    启动一个后台服务，该服务有如下两个作用：
    1. 监听`unix socket`，接受`sofa`发送过来到数据
    2. 监听 `redis` 任务队列，执行并处理结果

## build

    make clean && make build 

or

    go build -o ./bin/sofa ./main.go

## 


## run
       docker run -d -p 6379:6379 --name=redis redis
       sofa server -c ./sofa.yaml

## usage


### 1. start server

    sofa server -c ./sofa.yaml

### 2. collect data

   - 使用管道符收集收据 

    bjobs -W | sofa --platform=LSF --command=bjobs --tid=001


   - 使用exec直接执行命令获取数据
   
    sofa exec 'bjobs -W' -u test01 -t 001 -p lsf -c bjobs

### 3. 通过redis执行命令
可以向 `redis` `calls` 队列中发送消息，给`sofa server`来处理

消息结构：

    {
    	"TID": "001",
    	"Platform": "docker",
    	"Command": "info",
    	"FullCommand": "docker info"
    }

发送消息：
    
    lpush calls "{\"TID\":\"001\",\"Platform\":\"docker\",\"Command\":\"info\",\"FullCommand\":\"docker info\"}"
    
