# TextAutoCorrectionTool
This is a text autocorrection tool that reads text from a file and autocorrects the text and writes it in another text file.
Every instance of (hex) should replace the word before with the decimal version of the word.
Every instance of (bin) should replace the word before with the decimal version of the word 
Every instance of (up) converts the word before with the Uppercase version of it
Every instance of (low) converts the word before with the Lowercase version of it.
Every instance of (cap) converts the word before with the capitalized version of it.
or (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. 
Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".
The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. 
If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words 
Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h