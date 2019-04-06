// golangexample package containing 3 functions
package golangexamples

// including fmt for printing and greetings package to print greetings
import(
  "fmt"
  "github.com/ehteshamz/greetings"
)

//concatenate slices with a "-" in betwwen them
func ConcatSlice(sliceToConcat [] byte) string{
  // fmt.Printf(string(sliceToConcat))
  concatString:=" "
  for i:= range sliceToConcat {
    concatString+= (string(sliceToConcat[i]))
    concatString+="-"
  }
  return concatString[:(len(concatString)-1)]
}

//Encrypts each letter of given text with a letter some fixed number of positions down the alphabet
func Encrypt(sliceToEncrypt []byte,ceaserCount int){
  // fmt.Println(string(sliceToEncrypt))
  resultant :=" "
  for i:= range sliceToEncrypt {
    // x := (int(sliceToEncrypt[i])+ceaserCount-65)%26+65
    encrypted_letter := (int(sliceToEncrypt[i])+ceaserCount-65)%26+65
    resultant +=string(encrypted_letter)



    // if (string(sliceToEncrypt) ==string(sliceToEncrypt).toLowerCase() ){
    //   encrypted_letter := (int(sliceToEncrypt[i])+ceaserCount-97)%26+97
    //   resultant +=string(encrypted_letter)
    // }
    // if (string(sliceToEncrypt) ==string(sliceToEncrypt).toUpperCase() ){
    //   encrypted_letter := (int(sliceToEncrypt[i])+ceaserCount-65)%26+65
    //   resultant +=string(encrypted_letter)
    //
    // }


    // resultant += ((int(sliceToEncrypt[i])+ceaserCount-97)%26 +97);
  }
    fmt.Println(resultant)
}

func EzGreetings(name string)string{


  return greetings.PrintGreetings(name)
}
