package main

import "testing"

func BenchmarkTestString(b *testing.B){
	for n:= 0;n < b.N;n ++{
		testString()
	}
}


func BenchmarkTestStruct(b *testing.B){
	for n:= 0;n < b.N;n ++{
		testStruct()
	}
}


