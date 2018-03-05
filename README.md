# Anagram
A program for finding persian words anagrams written in go. This program has been tested on go version 1.9.3

## Overview
This program exports words from xlsx file, then sort each string and then add it to the hash map where the key is the sorted string and the value is an slice of strings where they are equal in sorted form. By doing this we can get each word's anagram with direct access ( O(1) ) by sorting the string and finding it values from the map.

## Installation
After downloading the project you need to set the GOPATH variable to the root of your project.

```export GOPATH=path/to/the/project```

After setting the GOPATH variable you have to get the package for reading xlsx files that i used in code. Type the following command : 

```go get github.com/tealeg/xlsx```

After installing the package you have to install the anagram package.

```go install anagram```

Then you have to run the compiled file from your $GOPATH/bin directory named "anagram".

## Usage
Select 1 from the menu ( Anagram ) and just type your word and you will get a slice of strings. 
