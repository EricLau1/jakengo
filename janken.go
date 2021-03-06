package jankengo

import(
  "fmt"
  "time"
  "math/rand"
  "errors"
)

var (
  ErrMatchDraw = errors.New("Match Draw!")
  ErrWrongChoice = errors.New("Wrong Choice!")
)

const (
  STONE   = 0
  PAPER   = 1
  SCISSOR = 2
)

var Itens = []string{"Stone", "Paper", "Scissor"}

type Player struct {
  Name string
  Choose int
}

func cpuToChoose() int {
  max := len(Itens)
  s := rand.NewSource(time.Now().UnixNano())
  r := rand.New(s)
  return r.Intn(max)
}

func Pow(player Player) (Player, error) {
  Init()
  cpu := Player{"CPU", cpuToChoose(),}
  winner, err := Start(player, cpu)
  if err != nil {
    return winner, err
  }
  return winner, nil
}

func Vs(p1, p2 Player) (Player, error) {
  Init()
  winner, err := Start(p1, p2)
  if err != nil {
    return winner, err
  }
  return winner, nil
}

func Start(p1, p2 Player) (Player, error) {
  if !isValidChoice(p1.Choose) {
    return Player{}, ErrWrongChoice
  }
  if verifyMatchDraw(p1, p2) {
    return Player{}, ErrMatchDraw
  }
  fmt.Printf("Player 1: %s ==> [%s]\n", p1.Name, Itens[p1.Choose])
  fmt.Printf("Player 2: %s ==> [%s]\n", p2.Name, Itens[p2.Choose])
  fmt.Println("=========================================")
  switch(p1.Choose) {
    case STONE:
      return toStone(p1, p2), nil
    case PAPER:
      return toPaper(p1, p2), nil
  }
  return toScissor(p1, p2), nil   
}

func isValidChoice(c int) bool {
  if c < 0 || c > len(Itens) - 1 {
    return false
  }
  return true
}

func verifyMatchDraw(p1, p2 Player) bool {
  if p1.Choose == STONE && p2.Choose == STONE {
    return true
  }
  if p1.Choose == PAPER && p2.Choose == PAPER {
    return true
  }
  if p1.Choose == SCISSOR && p2.Choose == SCISSOR {
    return true
  }
  return false
}

func toStone(p1, p2 Player) Player {
  if p1.Choose == STONE && p2.Choose == PAPER {
    return p2
  }
  return p1
}

func toPaper(p1, p2 Player) Player {
  if p1.Choose == PAPER && p2.Choose == SCISSOR {
    return p2
  }
  return p1
}

func toScissor(p1, p2 Player) Player {
  if p1.Choose == SCISSOR && p2.Choose == STONE {
    return p2
  }
  return p1
}

func Init() {
  var jankenpow = []string{"JAN", "-KEN-", "POW!!!"}
  for _, j := range jankenpow {
    fmt.Print(j)
    time.Sleep(1 * time.Second)
  }
  fmt.Println("\n")
}
