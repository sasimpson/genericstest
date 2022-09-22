# genericstest
some samples and tests for playing with generics. 


## example 1

i was playing with a client where the http methods were generic functions.  you could pass in the type of the thing you wanted returned and you would get the json data formatted and put into that type.  issue was adding some constraints, instead of having to add each type to a constraint file (might have been easier to understand and less work at this point), i made an interface to use as the type.  this meant i needed to add a method to really give the types a signature that would match the interface.  since i needed to unmarshal/decode some data into the structure, i made the decoder the method.

so the fun part comes from wanting to use that type to store the data, but also has the method which is writing back.  this did not go well unless its a pointer receiver.  if you add a * to make the type a pointer match, you're really just telling the compiler its a pointer to a type, not a type that can use a pointer. you can see the iteration in the tests and code of how that works out. 
