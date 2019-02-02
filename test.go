package jankengo

import(
  "fmt"
  "log"
)

func TestChooseStone() {
  p := Player{"Human", STONE,}
  winner, err := Pow(p)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Winner: %s, Choose: %s\n", winner.Name, Itens[winner.Choose])
}

func TestChoosePaper() {
  p := Player{"Human", PAPER,}
  winner, err := Pow(p)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("************************************")
  fmt.Printf("Winner: %s, Choose: %s\n", winner.Name,  Itens[winner.Choose])
  fmt.Println("************************************")
}

func TestChooseScissor() {
  p :=  Player{"Human",  SCISSOR,}
  winner, err :=  Pow(p)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("************************************")
  fmt.Printf("Winner: %s, Choose: %s\n", winner.Name,  Itens[winner.Choose])
  fmt.Println("************************************")
}

func TestWrongChoice() {
  p :=  Player{"Human", 3,} 
  _, err :=  Pow(p)
  if err ==  ErrWrongChoice {
    log.Println(err)
  }
  p.Choose = -1
  _, err =  Pow(p)
  if err ==  ErrWrongChoice {
    log.Println(err)
  }
}

