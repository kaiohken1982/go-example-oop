## Object example

Gestione OOP in go

# Confusion 

type P *int
func (P) f() { /* ... */ } // compile error: invalid receiver type

The (*Point).ScaleBy method can be cal le d by providing a *Point re ceiver, like this:

r := &Point{1, 2}
r.ScaleBy(2)
fmt.Println(*r) // "{2, 4}"

or this:

p := Point{1, 2}
pptr := &p
pptr.ScaleBy(2)
fmt.Println(p) // "{2, 4}"

or this:

p := Point{1, 2}
(&p).ScaleBy(2)
fmt.Println(p) // "{2, 4}"

the last two cas es are ungainly ( sgraziati )

# Utilizzo

Compilare

```
go build
```