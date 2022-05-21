package main

//Go supports embedding of structs and interfaces to
//express a more seamless composition of types.

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {

	return fmt.Sprintf("base with num=%v", b.num)

}

//A container embeds a base. An embedding looks like a field without a name.

type container struct {
	base

	str string
}

func main() {

	//When creating structs with literals, we have to initialize the embedding
	// explicitly; here the embedded type serves as the field name.

	co := container{

		base: base{

			num: 1,
		},

		str: "some name",
	}

	//We can access the base’s fields directly on co

	//Alternatively, we can spell out the full path using the embedded type name.

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	//Since container embeds base, the methods of base also
	//become methods of a container. Here we invoke a method
	//that was embedded from base directly on co.

	type describer interface {
		describe() string
	}

	var d describer = co

	fmt.Println("describer:", d.describe())

}

//Embedding structs with methods may be used to bestow interface implementations
//onto other structs. Here we see that a container now implements the describer interface because it embeds base.

//output :
//co={num: 1, str: some name}
//also num: 1
//describe: base with num=1
//describer: base with num=1
