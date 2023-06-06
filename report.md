# Wiggumize Report

__Scope:__
- https://0a0b00b404822135828b51d200ed008c.web-security-academy.net:443
- https://0af1000403f5a14b81ce752a009d0095.web-security-academy.net:443
- https://0abe003203cf09d58380418200940058.web-security-academy.net:443


__List of Checks:__
- __Secrets:__ This module is searching for secrets (eg. API keys)
- __LFI:__ This module is searching for filenames in request parameters. Could be an indication of possible LFI
- __SSRF:__ This module is searching for URL in request parameters.
- __notFound:__ This module is searching for 404 messages form hostings.
- __XML:__ This module is searching for XML in request parameters
- __Redirects:__ This module is searching for Redirects.
- __Parameters:__ This module is parsing GET or POST (JSON) params
--------------------

## SSRF
> This module is searching for URL in request parameters.
### Finding 0. - URL in a parameter
__Host: https://0abe003203cf09d58380418200940058.web-security-academy.net:443__ 

_Evidens:_

```
stockApi=http://192.168.0.1:8080/product/stock/check?productId=2&storeId=2
```
_More Details:_

```
URL:https://0abe003203cf09d58380418200940058.web-security-academy.net/product/stock
```
## XML
> This module is searching for XML in request parameters
### Finding 0. - possible XML in a parameter
__Host: https://0a0b00b404822135828b51d200ed008c.web-security-academy.net:443__ 

_Evidens:_

```
<?xml version="1.0" encoding="UTF-8"?><stockCheck><productId>2</productId><storeId>2</storeId></stockCheck>
```
_More Details:_

```
URL:https://0a0b00b404822135828b51d200ed008c.web-security-academy.net/product/stock
```
## Redirects
> This module is searching for Redirects.
### Finding 0. - Redirect Found
__Host: https://0af1000403f5a14b81ce752a009d0095.web-security-academy.net:443__ 

_Evidens:_

```
HTTP/2 302 Found
Location: /my-account
Set-Cookie: session=agd7bAvVyRgXZFPrhAGS3tTw4Yaz9gry; Secure; HttpOnly; SameSite=None
X-Frame-Options: SAMEORIGIN
Content-Length: 0


```
_More Details:_

```
Req Parameters:{"csrf":"oONNaOIWI5JJSLOvafTMskAuktH7Wyoc","username":"wiener","password":"peter"}
```
### Finding 1. - Redirect Found
__Host: https://0af1000403f5a14b81ce752a009d0095.web-security-academy.net:443__ 

_Evidens:_

```
HTTP/2 302 Found
Location: /login
X-Frame-Options: SAMEORIGIN
Content-Length: 0


```
## LFI
> This module is searching for filenames in request parameters. Could be an indication of possible LFI
### Finding 0. - filename in a parameter
__Host: https://0a0b00b404822135828b51d200ed008c.web-security-academy.net:443__ 

_Evidens:_

```
productId=test.php
```
_More Details:_

```
URL: https://0a0b00b404822135828b51d200ed008c.web-security-academy.net/product?productId=test.php
```


--------------------

## Parameters: 
__Host: https://0abe003203cf09d58380418200940058.web-security-academy.net:443__
_Endpoint: /product?productId=2_ 
- _productId:_ 2



__Host: https://0af1000403f5a14b81ce752a009d0095.web-security-academy.net:443__
_Endpoint: /product?productId=3_ 
- _productId:_ 3

_Endpoint: /product?productId=18_ 
- _productId:_ 18



__Host: https://0a0b00b404822135828b51d200ed008c.web-security-academy.net:443__
_Endpoint: /product?productId=test.php_ 
- _productId:_ test.php

_Endpoint: /product?productId=2_ 
- _productId:_ 2



