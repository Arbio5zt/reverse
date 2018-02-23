package reverse

import (
	"encoding/hex"
	"fmt"
)

func reverser(str string)string{
	//str :=[]byte{117, 58, 126, 81, 186, 44, 17, 169, 23, 103, 7, 131, 34, 8, 238, 70, 147, 33, 101, 140}
	//str:="0d821bd7b6d53f5c2b40e217c6defc8bbe896cf5"
	return ToHexString(ToArrayReverse((StringToByte(str))))
}

func ToArrayReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0 ;i--{
		x = append(x, arr[i])
	}
	return x
}

func ToHexString(data []byte) string {
	return hex.EncodeToString(data)
}

func StringToByte(str string) []byte{
	strx,_:= hex.DecodeString(str)
	return strx
}

