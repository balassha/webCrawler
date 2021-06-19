# Design Considerations
This application is designed as a Restful HTTP API service which can take an URL of the format
```json
<https://url> or <http://url>
```
I have built the application with Solid principles in mind. The concerns are layered and errors have been handled in all layers.

I have defined an interface (Parser) to abstract the implementation. This enables the application to be flexible in nature e.g. If we want to use a 3rd party library in future
we can make changes to the scrapper layer without modifying Parser.

I have used the concurrency pattern to identify the Inaccessible links. It enables us to improve the overall performance. 

I made a Trade-off to check the accessibility of a URL, I first wanted to send a HEAD HTTP request instead of GET. Because HEAD doesn't provide a response body and so should be faster.
But in my experience, I have seen web servers blocking HEAD requests with a 405 because it requires HTTPS by default. So I went ahead with the GET request. I chose reliability over speed which often times we as engineers are required to do.

The last requirement to identify if the webpage is a login page is a bit complicated. We can use certain things, such as checking the presense of a Login or signin button or check for input html tags with attributes of names username and password but all of this condition would match for a registration page as well. It is achievable but it requires few iterations to create a perfect solution. So in this application , I have created a subsidised module that checks for 'form' tags.

I have set a timeout of 10 seconds to check if the URL is Inaccessible. This can be configured based on the load.

I have also disabled SSL Verify in the HTTP client. It is to support websites with Self-Signed certificates.
