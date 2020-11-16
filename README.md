
# Hatch studio - Task

Write a small program in Go that accepts 2 json files and prints out if they are equal.

## Notes

1. json files could consist of a single json object or a list of json objects.

2. sorting of the json objects or the keys inside the json object does not matter for equality.

3. you have the freedom to use any pkg if any and structure it as you wish.

4. sizes of input files can vary from couple of lines to gigabytes

## Example of equal jsons

1. 

```
[{
	"id":"jhasdad",
	"name":"test json"
},
{
	"id":"wqweq",
	"name":"test json 2"
}]
```

2. 
```
[{
	"id":"wqweq",
	"name":"test json 2"
},
{
	"name":"test json",
	"id":"jhasdad"
}]
```

## Conditions

1. Please share the result via GitHub.

2. It’s a good idea to have a readme file. Please also document your decisions.

3. We appreciate it if you send the result within a week.

4. Have fun! :)

## Solution
#### Something to keep in mind:
1. Files can be big, composed of a lot of small object or only one huge object with several nested levels
2. At least one of the two file has to be all in memory to be compared with chunks of the second file
My idea is hash json objects in order to keep in memory small rappresentaions of it. 
Sha1 rappresentation of json fine already doesn't take in cosideration oreder of keys.
It is quite difficult have collitions, but it is still possible, we should consider it only if this start to be a problem, we could combine sha1 with another algorith to reduce this risk.
Garbage collector will clean memory after hashed inputs
goroutine will help to read two files quicker

during the comparison phase for each found we will remove those keys from source and target maps

* clone repository:
```
    git clone git@github.com:frenz/Hatch-task.git
    cd !$
```
* check help
```
    go run compare-json -h
```
by default is comparing input1.json and input2.json already added in the folder ./data
```
    go run compare-json -src-json<source file> -tgt-json<target file>
```
