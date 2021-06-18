#Design
This application is designed as a http api service which can take an URL of the format
<http/https://hostname>.

My Initial consideration was to build a WebUI along with Webserver to host the application. But later I decided to implement only the API service. The reason being,
It is a good practice to try and solve the core problem (It is the requirements in this case) rather than thinking about the surrounding areas.

I have built the application with Solid principles in mind. The concerns are layered and errors have been handled in all layers.

I have defined an interface to abstract the implementation and also to add support for using a separate 3rd party library instead of the webscrapper library that I created for this task.

The last requirement to identify if the webpage is a login page is a bit complicated, there is no clear way to identify if the page is a login page. We can use certain things, such as checking the presense of a Login or signin button or check for input html tags with attributes of names username and password but all of this condition would match for a registration page as well. It is achievable but it requires few iterations to create a perfect solution. SO in this application , I have created a subsidised check for 'form' tags.

I have set a timeout of 10 seconds to check if the URL is Inaccessible. This can be configured based on the load as 10 seconds is ideal.

