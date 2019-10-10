# sofa

## build

```shell script
 make clean && make build 
```
or

```shell script
 go build -o ./bin/sofa ./main.go
```

## Usage


### 1. start server


    sofa server -c=~/sofa.yaml



### 2. collect data

   - 使用管道符收集收据 

    bjobs -W | sofa --platform=LSF --command=bjobs --tid=001


   - 使用exec直接执行命令获取数据
   
    sofa exec 'bjobs -W' -u test01 -t 001 -p lsf -c bjobs


