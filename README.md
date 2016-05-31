## Function
- Make random anagrams from your input

## Sample
- Try to input any one word

```
$ go run anagrams.go dog
dgo
gdo
dog
odg
ogd
god
```

- 'ocean' is generated from 'canoe'

```
$ go run anagrams.go canoe | grep ocean
ocean
```

- Unicode characters can be input because input is treated as rune

```
$ go run anagrams.go あとうかい | grep かとう
あいかとう
いかとうあ
かとういあ
いあかとう
あかとうい
かとうあい
```

## Note
- It may have some bugs, because it sometimes be frozen.
- You can not see all anagrams of a long word, because maximum number of trials is 65536.
