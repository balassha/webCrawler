# Web Crawler
This application recevies a URL as an input and crawles the web page for the following
    - HTML Version
    - Page Title
    - Headings count by level
    - Amount of internal and external links
    - Amount of inaccessible links
    - If a page contains a login form

## Build & Run
'go build' will generate the binary of the name 'htmlparser'
Running the binary will launch the HTTP web server that listens on Port 8011. 

## Application & Usage
This application is a http api that hosts the following endpoint 
POST <url:8011>/parse/url

The body of the incoming request should be of the type down below
```josn
{
    "url":"https://www.google.com"
}
```

The response will be of the type given down below
```json
{
    "title": "Google",
    "version": "HTML 5",
    "links": {
        "ExternalLinks": [
            "https://www.google.co.in/imghp?hl=en&tab=wi",
            "https://maps.google.co.in/maps?hl=en&tab=wl",
            "https://play.google.com/?hl=en&tab=w8",
            "https://www.youtube.com/?gl=IN&tab=w1",
            "https://news.google.com/?tab=wn",
            "https://mail.google.com/mail/?tab=wm",
            "https://drive.google.com/?tab=wo",
            "https://www.google.co.in/intl/en/about/products?tab=wh",
            "http://www.google.co.in/history/optout?hl=en",
            "https://accounts.google.com/ServiceLogin?hl=en&passive=true&continue=https://www.google.com/&ec=GAZAAQ",
            "https://www.google.com/setprefs?sig=0_fBDKunrHij8Z9218glxgCGlyWKM%3D&hl=hi&source=homepage&sa=X&ved=0ahUKEwjE8qSCwaHxAhWzyzgGHYkgBmcQ2ZgBCAU",
            "https://www.google.com/setprefs?sig=0_fBDKunrHij8Z9218glxgCGlyWKM%3D&hl=bn&source=homepage&sa=X&ved=0ahUKEwjE8qSCwaHxAhWzyzgGHYkgBmcQ2ZgBCAY",
            "https://www.google.com/setprefs?sig=0_fBDKunrHij8Z9218glxgCGlyWKM%3D&hl=te&source=homepage&sa=X&ved=0ahUKEwjE8qSCwaHxAhWzyzgGHYkgBmcQ2ZgBCAc",
            "https://www.google.com/setprefs?sig=0_fBDKunrHij8Z9218glxgCGlyWKM%3D&hl=mr&source=homepage&sa=X&ved=0ahUKEwjE8qSCwaHxAhWzyzgGHYkgBmcQ2ZgBCAg",
            "https://www.google.com/setprefs?sig=0_fBDKunrHij8Z9218glxgCGlyWKM%3D&hl=ta&source=homepage&sa=X&ved=0ahUKEwjE8qSCwaHxAhWzyzgGHYkgBmcQ2ZgBCAk",
            "https://www.google.com/setprefs?sig=0_fBDKunrHij8Z9218glxgCGlyWKM%3D&hl=gu&source=homepage&sa=X&ved=0ahUKEwjE8qSCwaHxAhWzyzgGHYkgBmcQ2ZgBCAo",
            "https://www.google.com/setprefs?sig=0_fBDKunrHij8Z9218glxgCGlyWKM%3D&hl=kn&source=homepage&sa=X&ved=0ahUKEwjE8qSCwaHxAhWzyzgGHYkgBmcQ2ZgBCAs",
            "https://www.google.com/setprefs?sig=0_fBDKunrHij8Z9218glxgCGlyWKM%3D&hl=ml&source=homepage&sa=X&ved=0ahUKEwjE8qSCwaHxAhWzyzgGHYkgBmcQ2ZgBCAw",
            "https://www.google.com/setprefs?sig=0_fBDKunrHij8Z9218glxgCGlyWKM%3D&hl=pa&source=homepage&sa=X&ved=0ahUKEwjE8qSCwaHxAhWzyzgGHYkgBmcQ2ZgBCA0",
            "http://www.google.co.in/services/",
            "https://www.google.com/setprefdomain?prefdom=IN&prev=https://www.google.co.in/&sig=K_wZVpjI-qkMdRYbNE3kQaP9OtH1Q%3D"
        ],
        "InternalLinks": [
            "/preferences?hl=en",
            "/advanced_search?hl=en-IN&authuser=0",
            "/intl/en/ads/",
            "/intl/en/about.html",
            "/intl/en/policies/privacy/",
            "/intl/en/policies/terms/"
        ]
    },
    "headingsCount": null,
    "inaccessibleLinks": [
        "https://www.youtube.com/?gl=IN&tab=w1",
        "https://maps.google.co.in/maps?hl=en&tab=wl"
    ],
    "isLoginForm": true,
}
```
