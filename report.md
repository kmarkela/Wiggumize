# Wiggumize Report

__Scope:__
- https://0a1f009b032882b581481c2200e700a4.web-security-academy.net:443
- https://0afb0018038b33d080afdaa3005200db.web-security-academy.net:443


__List of Checks:__
- __Secrets:__ This module is searching for secrets (eg. API keys)
- __LFI:__ This module is searching for filenames in request parameters. Could be an indication of possible LFI
- __SSRF:__ This module is searching for URL in request parameters.
- __notFound:__ This module is searching for 404 messages form hostings.
- __XML:__ This module is searching for XML in request parameters
--------------------

## XML
> This module is searching for XML in request parameters
### Finding 0. - possible XML in a parameter
__Host: https://0a1f009b032882b581481c2200e700a4.web-security-academy.net:443__ 

_Evidens:_

```
 postId=<testing456>
```
_More Details:_

```
URL:https://0a1f009b032882b581481c2200e700a4.web-security-academy.net/post/comment/confirmation?postId=%3Ctesting456%3E
```
### Finding 1. - possible XML in a parameter
__Host: https://0a1f009b032882b581481c2200e700a4.web-security-academy.net:443__ 

_Evidens:_

```
 ------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="csrf" zILKDoLbGX7fVCaJ7NBsD83N3floBoqR
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="postId" 2
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="comment" test
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="name" pwn123
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="avatar"; filename="alt-battery-0-svgrepo-com.svg"
Content-Type: image/svg xml <?xml version="1.0" encoding="utf-8"?>
<!-- Uploaded to: SVG Repo, www.svgrepo.com, Generator: SVG Repo Mixer Tools -->
<svg fill="#000000" width="800px" height="800px" viewBox="0 0 32 32" version="1.1" xmlns="http://www.w3.org/2000/svg">
<title>alt-battery-1</title>
<path d="M0 20q0 2.496 1.76 4.256t4.256 1.76h17.984q2.496 0 4.256-1.76t1.76-4.256h1.984v-8h-1.984q0-2.464-1.76-4.224t-4.256-1.76h-17.984q-2.496 0-4.256 1.76t-1.76 4.224v8zM4 20v-8q0-0.832 0.576-1.408t1.44-0.576h17.984q0.832 0 1.408 0.576t0.608 1.408v8q0 0.832-0.608 1.44t-1.408 0.576h-17.984q-0.832 0-1.44-0.576t-0.576-1.44zM6.016 20h1.984v-8h-1.984v8z"></path>
</svg>
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="email" asd@ad.com
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="website" 
------WebKitFormBoundarybHN9oPufhQ35WXQD--

```
_More Details:_

```
URL:https://0a1f009b032882b581481c2200e700a4.web-security-academy.net/post/comment
```
## SSRF
> This module is searching for URL in request parameters.
### Finding 0. - URL in a parameter
__Host: https://0afb0018038b33d080afdaa3005200db.web-security-academy.net:443__ 

_Evidens:_

```
 postId=https://test.com
```
_More Details:_

```
URL:https://0afb0018038b33d080afdaa3005200db.web-security-academy.net/post?postId=https://test.com
```
### Finding 1. - URL in a parameter
__Host: https://0a1f009b032882b581481c2200e700a4.web-security-academy.net:443__ 

_Evidens:_

```
 ------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="csrf" zILKDoLbGX7fVCaJ7NBsD83N3floBoqR
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="postId" 2
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="comment" test
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="name" pwn123
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="avatar"; filename="alt-battery-0-svgrepo-com.svg"
Content-Type: image/svg xml <?xml version="1.0" encoding="utf-8"?>
<!-- Uploaded to: SVG Repo, www.svgrepo.com, Generator: SVG Repo Mixer Tools -->
<svg fill="#000000" width="800px" height="800px" viewBox="0 0 32 32" version="1.1" xmlns="http://www.w3.org/2000/svg">
<title>alt-battery-1</title>
<path d="M0 20q0 2.496 1.76 4.256t4.256 1.76h17.984q2.496 0 4.256-1.76t1.76-4.256h1.984v-8h-1.984q0-2.464-1.76-4.224t-4.256-1.76h-17.984q-2.496 0-4.256 1.76t-1.76 4.224v8zM4 20v-8q0-0.832 0.576-1.408t1.44-0.576h17.984q0.832 0 1.408 0.576t0.608 1.408v8q0 0.832-0.608 1.44t-1.408 0.576h-17.984q-0.832 0-1.44-0.576t-0.576-1.44zM6.016 20h1.984v-8h-1.984v8z"></path>
</svg>
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="email" asd@ad.com
------WebKitFormBoundarybHN9oPufhQ35WXQD
Content-Disposition: form-data; name="website" 
------WebKitFormBoundarybHN9oPufhQ35WXQD--

```
_More Details:_

```
URL:https://0a1f009b032882b581481c2200e700a4.web-security-academy.net/post/comment
```
### Finding 2. - URL in a parameter
__Host: https://0afb0018038b33d080afdaa3005200db.web-security-academy.net:443__ 

_Evidens:_

```
 username=http://127.0.0.1&password=asdasd
```
_More Details:_

```
URL:https://0afb0018038b33d080afdaa3005200db.web-security-academy.net/login
```
