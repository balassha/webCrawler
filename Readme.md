# Web Crawler
This application recevies a URL as an input, sends a HTTP Get to the URL and crawles the response web page for the following data
```
    - HTML Version
    - Page Title
    - Headings count by level
    - Amount of internal and external links
    - Amount of inaccessible links
    - If a page contains a login form
```
## Build & Run
'go build' will generate the binary of the name 'htmlparser'.
Running the binary will launch the HTTP web server that listens on Port 8011. 

## Application & Usage
This application is a Restful HTTP API that hosts the following endpoint 
```
POST <host:8011>/parse/url
```
The API accepts a JSON body. The structure of the input json is given down below.
```josn
{
    "url":"https://www.home24.com/websites/homevierundzwanzig/English/0/home24.html"
}
```
The application waits for 10 seconds to confirm if the External URL is inaccessible.
This is defined by the Timeout attribute in the HTTP client. It is configurable.

The response is of type JSON and a sample response is given down below
```yaml
{
    "title": "home24 | The online destination for home and living.",
    "version": "HTML 5",
    "headingsCount": {
        "h1": 0,
        "h2": 0,
        "h3": 3,
        "h4": 4,
        "h5": 0,
        "h6": 0
    },
    "links": {
        "ExternalLinks": [
            "https://www.home24.com/websites/homevierundzwanzig/English/1/homepage.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/7000/contact.html",
            "https://www.home24.com/websites/homevierundzwanzig/German/0/home24.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1000/about-us.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1100/who-we-are.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1200/story-of-home24.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1300/products-and-brands.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1400/a-unique-model.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/2000/our-team.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/2100/management-board.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/2200/supervisory-board.html",
            "https://home24.career.softgarden.de/en/",
            "https://www.home24.com/websites/homevierundzwanzig/English/4000/investor-relations.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4050/share.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4100/capital-increase.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4300/publications.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4400/corporate-governance.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4450/general-meeting.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4500/news.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4550/financial-calendar.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4600/ir-contact.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/5000/newsroom.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/5100/press-releases.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/6000/imprint.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/8000/data-protection-declaration.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/0/home24.html",
            "https://home24.career.softgarden.de/en/ ",
            "https://www.home24.de/",
            "https://www.home24.at/",
            "https://www.home24.ch/",
            "https://www.home24.be/",
            "https://www.home24.nl/",
            "https://www.home24.fr/",
            "https://www.home24.it/",
            "https://www.mobly.com.br/ ",
            "https://twitter.com/home24_de",
            "https://www.facebook.com/home24.de",
            "https://www.linkedin.com/company/home24/",
            "https://www.instagram.com/home24_de/",
            "https://www.youtube.com/user/home24TV",
            "https://www.pinterest.de/home24de/"
        ],
        "InternalLinks": [
            "javascript:void(0);",
            "tel:+493020389966",
            "mailto:info@home24.de",
            "tel:+4930700149000",
            "tel:+4930609880019"
        ]
    },
    "inaccessibleLinks": [
        "https://www.linkedin.com/company/home24/",
        "https://www.instagram.com/home24_de/",
        "https://home24.career.softgarden.de/en/ ",
        "https://www.mobly.com.br/ "
    ],
    "isLoginForm": true
}
```

## Another Example - https://www.home24.com/websites/homevierundzwanzig/English/8000/data-protection-declaration.html

```yaml
{   
    "title": "Data Protection Declaration | home24 Corporate Website",
    "version": "HTML 5",
    "headingsCount": {
        "h1": 0,
        "h2": 0,
        "h3": 13,
        "h4": 4,
        "h5": 1,
        "h6": 0
    },
    "links": {
        "ExternalLinks": [
            "https://www.home24.com/websites/homevierundzwanzig/English/1/homepage.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/7000/contact.html",
            "https://www.home24.com/websites/homevierundzwanzig/German/8000/data-protection-declaration.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1000/about-us.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1100/who-we-are.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1200/story-of-home24.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1300/products-and-brands.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/1400/a-unique-model.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/2000/our-team.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/2100/management-board.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/2200/supervisory-board.html",
            "https://home24.career.softgarden.de/en/",
            "https://www.home24.com/websites/homevierundzwanzig/English/4000/investor-relations.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4050/share.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4100/capital-increase.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4300/publications.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4400/corporate-governance.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4450/general-meeting.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4500/news.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4550/financial-calendar.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/4600/ir-contact.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/5000/newsroom.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/5100/press-releases.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/6000/imprint.html",
            "https://www.home24.com/websites/homevierundzwanzig/English/8000/data-protection-declaration.html",
            "https://www.home24.de/",
            "https://www.home24.at/",
            "https://www.home24.ch/",
            "https://www.home24.be/",
            "https://www.home24.nl/",
            "https://www.home24.fr/",
            "https://www.home24.it/",
            "https://www.mobly.com.br/ ",
            "https://twitter.com/home24_de",
            "https://www.facebook.com/home24.de",
            "https://www.linkedin.com/company/home24/",
            "https://www.instagram.com/home24_de/",
            "https://www.youtube.com/user/home24TV",
            "https://www.pinterest.de/home24de/"
        ],
        "InternalLinks": [
            "javascript:void(0);",
            "tel:+493020389966",
            "mailto:info@home24.de",
            "#_Toc14706386",
            "#_Toc14706386a",
            "#_Toc14706387a",
            "#_Toc14706387",
            "#_Toc14706388",
            "#_Toc14706389",
            "#_Toc14706390",
            "#_Toc14706391",
            "#_Toc14706392",
            "#_Toc14706393",
            "#_Toc14706394",
            "#_Toc14706395",
            "tel:+4930700149000",
            "tel:+4930609880019"
        ]
    },
    "inaccessibleLinks": [
        "https://www.instagram.com/home24_de/",
        "https://www.linkedin.com/company/home24/",
        "https://www.mobly.com.br/ "
    ],
    "isLoginForm": true
}
```
