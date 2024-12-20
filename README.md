# learning_go
## Exercises in the Learning Go by Jon Bodner ebook

While learning go through the Learning Go ebook, there are three different exercises on Page 113.

### Page 113:

1. Write a program the defines a variable named *greetings* of type slice of strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт". Create a subslice containing the first two values; a second subslice with the second, third and fourth values; and a third subslice with the fourth and fifth values. Print out all four slices.
2. Write a program that defines a string variable called *message* with the value "Hi 👩 and 👨" and prints the fourth rune in it as a character, not a number.
3. Write a program that defines a struct *Employee* with three fields: *firstName, lastName* and *id*. The first two fields are type of *string*, and the last field (*id*) is of type *int*. Create three instances of this struct using whatever values you'd like. Initialize the first one using the struct literal style without names, the second using the struct literal style with names, and the third with a *var* declaration. Use dot notation to populate the fields in the third struct. Print out all three structs.


### Chapter 4 exercises
1. Write a *for* loop that puts 100 random numbers betwee 0 and 100 into an *int* slice
2. Loop over the slice you created in exercsie 1. For each value in the slice, apply the following rules:
    - If the value is divisible by 2, print "Two!"
    - If the value is divisible by 3, print "Three!"
    - If the value is divisible by 2 and 3, print "Six!". Don't print anything else
    - Otherwise print "Never mind"
3. In *main*, declare and *int* variable called *total*. Write a *for* loop that uses a variable named *i* to iterate from 0 (inclusive) to 10 (exclusive). The body of the *for* loop should be as follows:
`total := total + i`
`fmt.Println(total)`
After the *for* loop, print out the value of *total*. What is printed out? What is the likely bug in this code?
