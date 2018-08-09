# password-generator-api
API for interaction generating passwords

## Description
This is an API that will generate passwords.   
The passwords will be alphanumerical containing letters, numbers and special characters.  
There is only one endpoint which is a GET endpoint that will provide multiple passwords.   
The endpoint is defined as ```/password/{min_length}/{no_of_numerical}/{no_of_symbols}/{no_of_passwords}```.   
As you can see you have to provide the minimum length of the generated passwords, the number of numerical characters,   
the number of special characters/symbols and the number of passwords you want to obtain.   

## Installation / Development
Just run ```make docker``` which will run a docker instance on your machine.   
Access the api through the docker ip.   

## Testing 
Just run ```make test``` to run the unit tests.   

## Code explanation

### main.go
The main.go file has is only starting the server.  
The routes definition and the route handlers are in a separate files so we can enable separation of concerns.  
Just imagine if you have an API with more than a 100 routes. This way it is simple to maintain the codebase  

### web/routes.go
The routes file is a small wrapper around the [mux](github.com/gorilla/mux) router.   
We provide a Routes structure:   
```
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
```
where the name, method, pattern and the handler function itself can be defined. It is completely flexible to extend.   

### web/handlers.go
The handlers file contains all the handlers needed for the API.   

#### getPassword handler
The main handler is the ```getPassword```. This is the handler that will request the password generation.  
If successful it will return a list of generated passwords.   

#### error handlers
The other handlers are the error handlers.   
The API provides custom handlers for Not found route (404), invalid request (422), and server error (500).   
All of them will return a json structure containing the error, type and message.   
An example would be:   
```
{
    "error":   "Validation Error",
    "type":    "422",
    "message": "The defined params are not valid or not allowed",
}
```

### components/helpers.go
The helpers file contains just a couple of helper functions we can use throughout the project.   
Currently it contains the ```ConvertToInt``` function that converts strings to integers   
and ```SetJsonHeader``` function that sets the http headers   


### generator/password_generator.go
The password_generator is the main part that actually generates the password.   
First we create 3 slices, one for letters, one for numbers and one for special characters/symbols.   
The length of the slices is defined by the user input through the http endpoint described earlier.   
For the `letters` slice we have a randomly conversion from vowels characters to numbers.    
When we convert a vowel to number we subtract the length of the `numbers` slice    
so we end up with exactly the same number of integers defined by the user.   
At the end we join all the slices in one resulting slice. After that the resulting slice is shuffled.   
The complexity of the algorithm is 3 O(n).   
The first time is when we iterate through the full length and we assign the values of the separate slices,   
the second time is when we do the appending/merging of the slices   
and the third time is when we shuffle the resulting slice.

