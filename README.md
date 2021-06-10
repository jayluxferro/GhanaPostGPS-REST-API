## GhanaPostGPS REST API
### Whitepaper
<p><a target="_blank" href="http://dx.doi.org/10.13140/RG.2.2.24355.27684/1">Read Now</a></p>

### API Details
<b>End Point URL:</b> https://ghanapostgps.sperixlabs.org<br>
<b>Method:</b> POST<br>
<b>Parameters:</b> address (GhanaPostGPS Address)<br>
<b>Content-Type:</b> application/x-www-form-urlencoded<br>
#### Output/Response:
1. Address found

```json
{
    "data": {
        "Table": [
            {
                "Area": "NEW KAGYASI",
                "CenterLatitude": 6.650080145273592,
                "CenterLongitude": -1.648700346667856,
                "District": "Kumasi",
                "EastLat": 6.65005768739201,
                "EastLong": -1.6486780409076,
                "GPSName": "AK4849321",
                "NorthLat": 6.65010262239948,
                "NorthLong": -1.6487229566718,
                "PostCode": "AK484",
                "Region": "Ashanti",
                "SouthLat": 6.65005768739201,
                "SouthLong": -1.6487229566718,
                "Street": "Kumasi, Ashanti, GHA",
                "WestLat": 6.65010262239948,
                "WestLong": -1.6486780409076
            }
        ]
    },
    "found": true
}
```

2. Address not found

```json
{
    "data": {
        "Table": null
    },
    "found": false
}
```

### Sample Codes
<a href="#csharp">C-Sharp</a> | <a href="#curl">cURL</a> | <a href="#go">Golang</a> | <a href="#js">Javascript</a> | <a href="#node">NodeJS</a> | <a href="#php">PHP</a> | <a href="#python">Python</a> | <a href="#swift">Swift</a> | <a href="#java">Java</a> | <a href="#ruby">Ruby</a> | <a href="#powershell">PowerShell</a><br><br>
<hr/>
<b>Address:</b> AK-484-9321 or AK4849321<br><br>

<hr id="csharp">
<h4>C-Sharp</h4>
<hr/>
Code:

```javascript
var client = new RestClient("https://ghanapostgps.sperixlabs.org");
client.Timeout = -1;
var request = new RestRequest(Method.POST);
request.AddHeader("Content-Type", "application/x-www-form-urlencoded");
request.AddParameter("address", "AK-484-9321");
request.OnBeforeDeserialization = resp => { resp.ContentType = "application/json"; };
IRestResponse response = client.Execute(request);
Console.WriteLine(response.Content);
```

<hr id="curl">
<h4>cURL</h4>
<hr/>
Code:

```bash
curl --location --request POST 'https://ghanapostgps.sperixlabs.org' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'address=AK-484-9321'
```

<hr id="go">
<h4>Go</h4>
<hr/>
Code:

```go
package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://ghanapostgps.sperixlabs.org"
  method := "POST"

  payload := strings.NewReader("address=AK-484-9321")

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}
```

<hr id="js">
<h4>Javscript</h4>
<hr/>
Code:

```javascript
var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

var urlencoded = new URLSearchParams();
urlencoded.append("address", "AK-484-9321");

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: urlencoded,
  redirect: 'follow'
};

fetch("https://ghanapostgps.sperixlabs.org", requestOptions)
  .then(response => response.json())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```



<hr id="node">
<h4>NodeJS</h4>
<hr/>
Code:

```javascript
var request = require('request');
var options = {
  'method': 'POST',
  'url': 'https://ghanapostgps.sperixlabs.org',
  'headers': {
    'Content-Type': 'application/x-www-form-urlencoded'
  },
  form: {
    'address': 'AK-484-9321'
  }
};
request(options, function (error, response) {
  if (error) throw new Error(error);
  console.log(response.body);
});
```


<hr id="php">
<h4>PHP</h4>
<hr/>
Code:

```php
<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => "https://ghanapostgps.sperixlabs.org",
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => "",
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => "POST",
  CURLOPT_POSTFIELDS => "address=AK-484-9321",
  CURLOPT_HTTPHEADER => array(
    "Content-Type: application/x-www-form-urlencoded"
  ),
));

$response = curl_exec($curl);

curl_close($curl);
echo $response;
```


<hr id="python">
<h4>Python</h4>
<hr/>
Code:

```python
import requests

url = "https://ghanapostgps.sperixlabs.org"

payload = 'address=AK-484-9321'
headers = {
  'Content-Type': 'application/x-www-form-urlencoded'
}

response = requests.request("POST", url, headers=headers, data = payload)

print(response.json())
```


<hr id="swift">
<h4>Swift</h4>
<hr/>
Code:

```swift
import Foundation

var semaphore = DispatchSemaphore (value: 0)

let parameters = "address=AK-484-9321"
let postData =  parameters.data(using: .utf8)

var request = URLRequest(url: URL(string: "https://ghanapostgps.sperixlabs.org")!,timeoutInterval: Double.infinity)
request.addValue("application/x-www-form-urlencoded", forHTTPHeaderField: "Content-Type")

request.httpMethod = "POST"
request.httpBody = postData

let task = URLSession.shared.dataTask(with: request) { data, response, error in 
  guard let data = data else {
    print(String(describing: error))
    return
  }
  print(String(data: data, encoding: .utf8)!)
  semaphore.signal()
}

task.resume()
semaphore.wait()
```


<hr id="java">
<h4>Java</h4>
<hr/>
Code:

```java
OkHttpClient client = new OkHttpClient().newBuilder()
  .build();
MediaType mediaType = MediaType.parse("application/x-www-form-urlencoded");
RequestBody body = RequestBody.create(mediaType, "address=AK-484-9321");
Request request = new Request.Builder()
  .url("https://ghanapostgps.sperixlabs.org")
  .method("POST", body)
  .addHeader("Content-Type", "application/x-www-form-urlencoded")
  .build();
Response response = client.newCall(request).execute();
```

<hr id="ruby">
<h4>Ruby</h4>
<hr/>
Code:

```ruby
require "uri"
require "net/http"

url = URI("https://ghanapostgps.sperixlabs.org")

https = Net::HTTP.new(url.host, url.port);
https.use_ssl = true

request = Net::HTTP::Post.new(url)
request["Content-Type"] = "application/x-www-form-urlencoded"
request.body = "address=AK-484-9321"

response = https.request(request)
puts response.read_body
```


<hr id="powershell">
<h4>PowerShell</h4>
<hr/>
Code:

```powershell
$headers = New-Object "System.Collections.Generic.Dictionary[[String],[String]]"
$headers.Add("Content-Type", "application/x-www-form-urlencoded")

$body = "address=AK-484-9321"

$response = Invoke-RestMethod 'https://ghanapostgps.sperixlabs.org' -Method 'POST' -Headers $headers -Body $body
$response | ConvertTo-Json
```

