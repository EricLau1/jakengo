package main

import(
  "fmt"
  "jakengo/janken"
  "log"
)

func main(){
  TestChooseStone()
  TestChoosePaper()
  TestChooseScissor()
  TestWrongChoice()
}

func TestChooseStone() {
  p := janken.Player{"Human", janken.STONE,}
  winner, err := janken.Pow(p)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("************************************")
  fmt.Printf("Winner: %s, Choose: %s\n", winner.Name, janken.Itens[winner.Choose])
  fmt.Println("************************************")
}

func TestChoosePaper() {
  p := janken.Player{"Human", janken.PAPER,}
  winner, err := janken.Pow(p)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("************************************")
  fmt.Printf("Winner: %s, Choose: %s\n", winner.Name, janken.Itens[winner.Choose])
  fmt.Println("************************************")
}

func TestChooseScissor() {
  p := janken.Player{"Human", janken.SCISSOR,}
  winner, err := janken.Pow(p)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("************************************")
  fmt.Printf("Winner: %s, Choose: %s\n", winner.Name, janken.Itens[winner.Choose])
  fmt.Println("************************************")
}

func TestWrongChoice() {
  p := janken.Player{"Human", 3,} 
  _, err := janken.Pow(p)
  if err == janken.ErrWrongChoice {
    log.Println(err)
  }
  p.Choose = -1
  _, err = janken.Pow(p)
  if err == janken.ErrWrongChoice {
    log.Println(err)
  }
}
