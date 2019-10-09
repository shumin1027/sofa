# sofa

## build
```shell script
 go build -mod=vendor -o ./bin/sofa 
```



## Usage


### 1. start server


    sofa server -c=~/sofa.yaml



### 2. collect data


    bjobs -W | sofa --platform=LSF --command=bjobs --tid=001

or

    sofa exec 'bjobs -W' --platform=LSF --command=bjobs --tid=001


