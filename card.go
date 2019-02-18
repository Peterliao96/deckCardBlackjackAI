package deckCardBlackjackAI

type Suit uint8

const (
  Spade Suit = iota // increment by 1
  Diamond
  Club
  Heart
  Joker
)

type Rank uint8

const (
  _ Rank = iota
  Ace
  Two
  Three
  Four
  Five
  Six
  Seven
  Eight
  Nine
  Ten
  Jack
  Queen
  King
)

type Card struct {
  Suit
  Rank
}