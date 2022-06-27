## GhanaPostGPS REST API
### Whitepaper
<p><a target="_blank" href="http://dx.doi.org/10.13140/RG.2.2.24355.27684/2">Read Now</a></p>


<details>
<summary>Deploy your own instance</summary>
<hr/>

### Local Server / VPS
1. Download and install Golang (<a href="https://golang.org/dl" target="_blank">here</a>). Make sure it's added to your environment path.

2. Clone the repository.
```
git clone https://github.com/jayluxferro/GhanaPostGPS-REST-API.git ghanapostgps
```

3. Install dependencies.
```terminal
cd ghanapostgps
go mod download
go mod vendor
```

4. Run in development mode
```
./dev
```

5. Run in live mode
```
./live
```

**NB:** 
* Default port is `5001`. Modify the 'run' scripts to change the default port. The API documentation is the same; replace the hostname with your instance's.

* The default API keys are stored in the `.env` file. You can generate new keys <a href="https://ghanapostgps.sperixlabs.org" target="_blank">here</a>.

### Heroku
1. Clone the project.

2. Create a new heroku project and add its git URL. Example is shown below:
```
    git remote add heroku https://myproject.heroku.git
```

3. Push to the heroku instance.
```
    git push heroku master
```
</details>
<br/>


<details>
<summary>Get Location Coordinates (Latitude/Longitude) from GhanaPostGPS Address.</summary>
<hr/>
### API Details<br>
<b>End Point URL:</b> https://ghanapostgps.sperixlabs.org/get-location<br>
<b>Method:</b> POST<br>
<b>API Parameters:</b> address (GhanaPostGPS Address)<br>
<b>Content-Type:</b> application/x-www-form-urlencoded<br><br>
#### Output/Response:<br>

1. Location found

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

2. Location not found

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
var client = new RestClient("https://ghanapostgps.sperixlabs.org/get-location");
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
curl --location --request POST 'https://ghanapostgps.sperixlabs.org/get-location' \
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

  url := "https://ghanapostgps.sperixlabs.org/get-location"
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

fetch("https://ghanapostgps.sperixlabs.org/get-location", requestOptions)
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
  'url': 'https://ghanapostgps.sperixlabs.org/get-location',
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
  CURLOPT_URL => "https://ghanapostgps.sperixlabs.org/get-location",
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

url = "https://ghanapostgps.sperixlabs.org/get-location"

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

var request = URLRequest(url: URL(string: "https://ghanapostgps.sperixlabs.org/get-location")!,timeoutInterval: Double.infinity)
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
  .url("https://ghanapostgps.sperixlabs.org/get-location")
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

url = URI("https://ghanapostgps.sperixlabs.org/get-location")

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

$response = Invoke-RestMethod 'https://ghanapostgps.sperixlabs.org/get-location' -Method 'POST' -Headers $headers -Body $body
$response | ConvertTo-Json
```
</details>
<br>


<details>
<summary>Get GhanaPostGPS Address from Location Coordinates (Latitude/Longitude).</summary>
<hr/>
### API Details<br>
<b>End Point URL:</b> https://ghanapostgps.sperixlabs.org/get-address<br>
<b>Method:</b> POST<br>
<b>API Parameters:</b> lat (latitude), long (longitude)<br>
<b>Content-Type:</b> application/x-www-form-urlencoded<br><br>
#### Output/Response:<br>

1. Address found

```json
{
    "data": {
      "Table": [
        {
          "GPSName": "AK4849319",
          "Region": "Ashanti",
          "District": "Kumasi",
          "PostCode": "AK484",
          "NLat": 6.650012752389040,
          "SLat": 6.649967817390580,
          "WLong": -1.648722956671800,
          "Elong": -1.648678040907600,
          "Area": "NEW KAGYASI",
          "Street": "Kumasi, Ashanti, GHA"
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
<a href="#csharp2">C-Sharp</a> | <a href="#curl2">cURL</a> | <a href="#go2">Golang</a> | <a href="#js2">Javascript</a> | <a href="#node2">NodeJS</a> | <a href="#php2">PHP</a> | <a href="#python2">Python</a> | <a href="#swift2">Swift</a> | <a href="#java2">Java</a> | <a href="#ruby2">Ruby</a> | <a href="#powershell2">PowerShell</a><br><br>
<hr/>
<b>Address:</b> AK-484-9321 or AK4849321<br><br>

<hr id="csharp2">
<h4>C-Sharp</h4>
<hr/>
Code:

```javascript
var client = new RestClient("https://ghanapostgps.sperixlabs.org/get-address");
client.Timeout = -1;
var request = new RestRequest(Method.POST);
request.AddHeader("Content-Type", "application/x-www-form-urlencoded");
request.AddParameter("lat", "6.6500");
request.AddParameter("long", "-1.6487");
request.OnBeforeDeserialization = resp => { resp.ContentType = "application/json"; };
IRestResponse response = client.Execute(request);
Console.WriteLine(response.Content);
```

<hr id="curl2">
<h4>cURL</h4>
<hr/>
Code:

```bash
curl --location --request POST 'https://ghanapostgps.sperixlabs.org/get-address' --form 'lat="6.6500"' --form 'long="-1.647"'
```

<hr id="go2">
<h4>Go</h4>
<hr/>
Code:

```go
package main

import (
  "fmt"
  "bytes"
  "mime/multipart"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://ghanapostgps.sperixlabs.org/get-address"
  method := "POST"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  _ = writer.WriteField("lat", "6.6500")
  _ = writer.WriteField("long", "-1.647")
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
    return
  }


  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
```

<hr id="js2">
<h4>Javscript</h4>
<hr/>
Code:

```javascript
var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

var formdata = new FormData();
formdata.append("lat", "6.6500");
formdata.append("long", "-1.647");

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: formdata
};

fetch("https://ghanapostgps.sperixlabs.org/get-address", requestOptions)
  .then(response => response.json())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
```



<hr id="node2">
<h4>NodeJS</h4>
<hr/>
Code:

```javascript
var request = require('request');
var options = {
  'method': 'POST',
  'url': 'https://ghanapostgps.sperixlabs.org/get-address',
  'headers': {
    'Content-Type': 'application/x-www-form-urlencoded'
  },
  formData: {
    'lat': '6.6500',
    'long': '-1.647'
  }
};
request(options, function (error, response) {
  if (error) throw new Error(error);
  console.log(response.body);
});
```


<hr id="php2">
<h4>PHP</h4>
<hr/>
Code:

```php
<?php

$curl = curl_init();

curl_setopt_array($curl, array(
  CURLOPT_URL => 'https://ghanapostgps.sperixlabs.org/get-address',
  CURLOPT_RETURNTRANSFER => true,
  CURLOPT_ENCODING => '',
  CURLOPT_MAXREDIRS => 10,
  CURLOPT_TIMEOUT => 0,
  CURLOPT_FOLLOWLOCATION => true,
  CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
  CURLOPT_CUSTOMREQUEST => 'POST',
  CURLOPT_POSTFIELDS => http_build_query(array('lat' => '6.6500','long' => '-1.647')),
  CURLOPT_HTTPHEADER => array(
    'Content-Type: application/x-www-form-urlencoded'
  ),
));

$response = curl_exec($curl);

curl_close($curl);
echo $response;
```


<hr id="python2">
<h4>Python</h4>
<hr/>
Code:

```python
import requests

url = "https://ghanapostgps.sperixlabs.org/get-address"

payload={'lat': '6.6500', 'long': '-1.647'}
files=[]
headers = {
  'Content-Type': 'application/x-www-form-urlencoded'
}

response = requests.request("POST", url, headers=headers, data=payload, files=files)

print(response.json())
```


<hr id="swift2">
<h4>Swift</h4>
<hr/>
Code:

```swift
import Foundation
#if canImport(FoundationNetworking)
import FoundationNetworking
#endif

var semaphore = DispatchSemaphore (value: 0)

let parameters = [
  [
    "key": "lat",
    "value": "6.6500",
    "type": "text"
  ],
  [
    "key": "long",
    "value": "-1.647",
    "type": "text"
  ]] as [[String : Any]]

let boundary = "Boundary-\(UUID().uuidString)"
var body = ""
var error: Error? = nil
for param in parameters {
  if param["disabled"] == nil {
    let paramName = param["key"]!
    body += "--\(boundary)\r\n"
    body += "Content-Disposition:form-data; name=\"\(paramName)\""
    if param["contentType"] != nil {
      body += "\r\nContent-Type: \(param["contentType"] as! String)"
    }
    let paramType = param["type"] as! String
    if paramType == "text" {
      let paramValue = param["value"] as! String
      body += "\r\n\r\n\(paramValue)\r\n"
    } else {
      let paramSrc = param["src"] as! String
      let fileData = try NSData(contentsOfFile:paramSrc, options:[]) as Data
      let fileContent = String(data: fileData, encoding: .utf8)!
      body += "; filename=\"\(paramSrc)\"\r\n"
        + "Content-Type: \"content-type header\"\r\n\r\n\(fileContent)\r\n"
    }
  }
}
body += "--\(boundary)--\r\n";
let postData = body.data(using: .utf8)

var request = URLRequest(url: URL(string: "https://ghanapostgps.sperixlabs.org/get-address")!,timeoutInterval: Double.infinity)
request.addValue("application/x-www-form-urlencoded", forHTTPHeaderField: "Content-Type")
request.addValue("multipart/form-data; boundary=\(boundary)", forHTTPHeaderField: "Content-Type")

request.httpMethod = "POST"
request.httpBody = postData

let task = URLSession.shared.dataTask(with: request) { data, response, error in 
  guard let data = data else {
    print(String(describing: error))
    semaphore.signal()
    return
  }
  print(String(data: data, encoding: .utf8)!)
  semaphore.signal()
}

task.resume()
semaphore.wait()
```

<hr id="ruby2">
<h4>Ruby</h4>
<hr/>
Code:

```ruby
require "uri"
require "net/http"

url = URI("https://ghanapostgps.sperixlabs.org/get-address")

http = Net::HTTP.new(url.host, url.port);
request = Net::HTTP::Post.new(url)
request["Content-Type"] = "application/x-www-form-urlencoded"
form_data = [['lat', '6.6500'],['long', '-1.647']]
request.set_form form_data, 'multipart/form-data'
response = http.request(request)
puts response.read_body
```


<hr id="powershell2">
<h4>PowerShell</h4>
<hr/>
Code:

```powershell
$headers = New-Object "System.Collections.Generic.Dictionary[[String],[String]]"
$headers.Add("Content-Type", "application/x-www-form-urlencoded")

$multipartContent = [System.Net.Http.MultipartFormDataContent]::new()
$stringHeader = [System.Net.Http.Headers.ContentDispositionHeaderValue]::new("form-data")
$stringHeader.Name = "lat"
$stringContent = [System.Net.Http.StringContent]::new("6.6500")
$stringContent.Headers.ContentDisposition = $stringHeader
$multipartContent.Add($stringContent)

$stringHeader = [System.Net.Http.Headers.ContentDispositionHeaderValue]::new("form-data")
$stringHeader.Name = "long"
$stringContent = [System.Net.Http.StringContent]::new("-1.647")
$stringContent.Headers.ContentDisposition = $stringHeader
$multipartContent.Add($stringContent)

$body = $multipartContent

$response = Invoke-RestMethod 'https://ghanapostgps.sperixlabs.org/get-address' -Method 'POST' -Headers $headers -Body $body
$response | ConvertTo-Json
```
</details>
