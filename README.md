# http-header-checker

Log every HTTP headers in the request.

## Usage

```bash
-address string
    Listen address (default ":8080")

```
## Example

```bash
curl -H "TestHeader: test" http://localhost:8080/
```
```
New request from 127.0.0.1:40668
	User-Agent -> [curl/7.86.0]
	Accept -> [*/*]
	Testheader -> [test]
---------------
```