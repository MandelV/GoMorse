# GoMorse
![Go](https://github.com/MandelV/GoMorse/workflows/Go/badge.svg)

Go library that enable to Encode/Decode Morse
## Installation
```Go
import (
 "github.com/MandelV/GoMorse"
)

```
## Usage

There are some methodes to Encode message to morse and Decode morse to plaintext
or get the letter of a given code and code for given letter

### To Encode

```Go
func Encode(message *string) (morse *string, err error)
```
Encode take a pointer of the message you want to encode 

```
if Encode succeed : 
return *morse, nil
else :
return nil, err
```

### To Decode

```Go
func Decode(morse *string) (message *string, err error)
```
Decode take a pointer of the morse you want to decode 

```
if Decode succeed : 
return *message, nil
else :
return nil, err
```

### Get Letter of given morse code

```Go
func GetLetter(morse string) (letter string, err error)
```
Get the letter against its morse code
```
if GetLetter succeed : 
return "", nil
else :
return nil, err
```


### Get Code of given letter

```Go
func GetCode(letter string) (code string, err error)
```
Get the morse code against its letter
```
if GetCode succeed : 
return "", nil
else :
return nil, err
```

## Overview 

Internally the package represente the morse langage into Binary Tree :
```Go
//Tree represents the morse language
type Tree struct {
	Groot *Node
}
```
The tree contains nodes, each node represent a letter

```Go
//Node represent a letter in Tree
type Node struct {
	Dot    *Node
	Dash   *Node
	Letter string
}
```
To get the code of a letter you have to browse through the binary tree if the path goes left its a dot (.) otherwise its a dash (-)

Here a representation of the binary tree :

![img](doc/morse_tree.png)

## TO-DO

- [Add the ability to manage empty nodes in the tree](https://github.com/MandelV/GoMorse/issues/1)
