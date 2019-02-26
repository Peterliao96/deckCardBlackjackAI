package deck

import ("fmt"
"testing")

func ExampleCard(){
  fmt.Print(Card{Rank:Ace,Suit:Heart})

  // Output:
  // Ace of Heart
}

func TestNew(t *testing.T){
  cards := New()
  if len(cards) != 13*4{
    t.Error("wrong number of cards in a new deck")
  }
}
