Simple JAN-KEN-POW Game    

Download package 

    go get -u github.com/EricLau1/jankengo

Import package

    import(
        "github.com/EricLau1/jankengo"
    )


variable publishes with available options

    var Itens = []string{"Stone", "Paper", "Scissor"}

Player Struct 

    type Player struct {
        Name string
        Choose int
    }

Constants

    const (
        STONE   = 0
        PAPER   = 1
        SCISSOR = 2
    )

Errors
 
The match ended in a draw.
  
    ErrMatchDraw
  
Chose greater than 2 or less than 0

    ErrWrongChoice


How to play

    p := jankengo.Player{"Human", jankengo.STONE,}
    winner, err := jankengo.Pow(p)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Winner: %s, Choose: %s\n", winner.Name, jankengo.Itens[winner.Choose])

How to play Vs mode

    p1 := jankengo.Player{"Jane Doe", jankengo.STONE,}
    p2 := jankengo.Player{"Jon Doe", jankengo.SCISSOR,}
    winner, err := jankengo.Vs(p1, p2)
    if err != nil {
    log.Fatal(err)
    }
    fmt.Printf("Winner: %s ==> [%s]\n", winner.Name, jankengo.Itens[winner.Choose])
