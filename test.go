package jankengo

import(
  "fmt"
  "log"
)
func main(){
  TestChooseStone()
  TestChoosePaper()
  TestChooseScissor()
  TestWrongChoice()
}

func TestChooseStone() {
  p := jankengo.Player{"Human", jankengo.STONE,}
  winner, err := jankengo.Pow(p)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Winner: %s, Choose: %s\n", winner.Name, jankengo.Itens[winner.Choose])
}

func TestChoosePaper() {
  p := jankengo.Player{"Human", jankengo.PAPER,}
  winner, err := jankengo.Pow(p)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("************************************")
  fmt.Printf("Winner: %s, Choose: %s\n", winner.Name, jankengo.Itens[winner.Choose])
  fmt.Println("************************************")
}

func TestChooseScissor() {
  p := jankengo.Player{"Human", jankengo.SCISSOR,}
  winner, err := jankengo.Pow(p)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("************************************")
  fmt.Printf("Winner: %s, Choose: %s\n", winner.Name, jankengo.Itens[winner.Choose])
  fmt.Println("************************************")
}

func TestWrongChoice() {
  p := jankengo.Player{"Human", 3,} 
  _, err := jankengo.Pow(p)
  if err == jankengo.ErrWrongChoice {
    log.Println(err)
  }
  p.Choose = -1
  _, err = jankengo.Pow(p)
  if err == jankengo.ErrWrongChoice {
    log.Println(err)
  }
}

