struct Person {
    Name: String
}

func main() {
    var p = new Person { Name: "John" }
    println(p.Name) //output: John -> 1 object is created
    p.Name = "Doe"
    println(p.Name) //output: Doe -> 1 object is created
}