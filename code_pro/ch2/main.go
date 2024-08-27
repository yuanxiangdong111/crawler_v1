package main

type Fruit interface {
    show()
}

type Apple struct {
    Fruit
}

type Banana struct {
    Fruit
}

type Orange struct {
    Fruit
}

func (a *Apple) show() {
    println("I am an apple")
}

func (b *Banana) show() {
    println("I am an banana")
}

func (o *Orange) show() {
    println("I am an orange")
}

func main() {
    var f Fruit
    f = &Apple{}
    f.show()
    f = &Banana{}
    f.show()
    f = &Orange{}
    f.show()

}
