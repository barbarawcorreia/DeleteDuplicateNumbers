package main

import (
	"sort"
)

func main() {

}

//PROBLEMA: TIRAR DUPLICADAS DO ARRAY.
// Primeira tentativa Baby Steps, como se já estivesse ordenado - TestSlicesDuplicatedSequency
// func DeleteDuplicated(s []int) []int {
// 	sort.Ints(s)
// 	sliceUniqueNumbers := []int{}
// 	for index, _ := range s {
// 		if index < len(s)-1 && s[index] == s[index+1] {
// 			sliceUniqueNumbers = append(sliceUniqueNumbers, s[index+1])
// 		}
// 		fmt.Println(sliceUniqueNumbers)
// 	}
// 	return sliceUniqueNumbers
// }

// len(s)-1 é usado para que não haja o erro index out of range.
// que acaba acessando uma posição não existente.
// então para não ultrapassar o tamanho, iteramos somente até
// o tamanho do slice menos 1 posição.

// //2 Tentantiva, como se estivesse desordenado - TestSlicesDuplicatedNotInSequency
func DeleteDuplicated(s []int) []int {
	sort.Ints(s)
	indexOut := 0
	for indexInside := 1; indexInside < len(s); indexInside++ {
		if s[indexOut] == s[indexInside] {
			continue
		}
		indexOut++
		s[indexOut] = s[indexInside]
	}
	result := s[0 : indexOut+1]
	return result
}

//Lógica : compara a primeira posicao com todos, se for igual, volta pro loop
//ignora e continua até achar diferente.
// achando diferente, troca a posicao
// no retorno, imprime so os que foram filtrados, deixando os repetidos pro final.
