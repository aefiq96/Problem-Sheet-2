# Web Application
answer to go-problem-quick

# 1. Guessing game
Git clone my project 
go build my project 
execute the file
It should respond with the text “Guessing game” in the browser

# 2. Added a guess route 

# 3.Turned the guess page into a template
Change the web application to use the guess.tmpl file as a template. Add a single variable to the template called 
Message and create  a struct in Go containing a single string. Create an instance of the struct with the string set
to “Guess a number between 1 and 20” and have the template render this in the H2 tag.

# 4. Add a guess route

Create a new route in your application at /guess. Have it serve a new page called guess.html. Use the same Bootstrap code for the page as in index.html but add a level 2 heading with the text “Guess a number between 1 and 20”. Add a form, with a text input with the name “guess” and a submit button with the label “Guess”. The action of the form should be /guess and the method should be GET.

# 4. Turn the guess page into a template

Change the web application to use the guess.tmpl file as a template. Add a single variable to the template called Message and create a struct in Go containing a single string. Create an instance of the struct with the string set to “Guess a number between 1 and 20” and have the template render this in the H2 tag.

# 5. Check for a cookie

In the /guess handler check if the cookie called target is set. If it is not, generate a random number between 1 and 20 and set a cookie called target to that value in the response. Otherwise, leave the cookie at its current value.

# 6. Check for a variable

Have the /guess handler check if a URL encoded variable called guess has been set to an integer. If it has, have the text “You guessed {guess}.” inserted into the template where {guess} is replaced with the value of guess.

# 7. Compare the cookie and the guess

If the target cookie and the guess variable are both set, then have the /guess handler compare them. If they are equal, set the target cookie to another randomly generated integer, and have the template display a congratulations message and a link to create a new game. Otherwise, have the template display a message telling the user what their guess was and whether it was too high or too low.

# 8. Use the POST method

Change your HTML form in guess.html to use the POST method instead. Make sure your application still works, bug fixing it if necessary.
