# home-sso-service
This is Single-Sign On service
## Dependencies
- go version go1.15.6 linux/amd64

## Build

To build the project, use:

```bash
# Use 
go get .
```
## Run

To run the project, use:

```bash
# Use 
go run main.go
```
## API

Request OAuth2 JWT token

```bash
POST http://192.168.0.106:1323/api/sso/requestToken
Body
{
    "clientid":"exampleOauthClient",
    "clientsecret":"exampleOauthClientSecret",
    "username":"thanhlam",
    "password":"12345678"
}
```
Parse OAuth2 JWT token

```bash
POST http://192.168.0.106:1323/api/sso/parseToken
Body
{
    "token":"qFfbbFKLO3ArNEhSDiti4_zdjwysV7J1cWyHE84gOe7SDXJPIiQ_GWW2VEULaqwYpLXu_RR5PHzXfnY9rDWo7_2PGNzJfCLihU_aTfX42XkjOFS8HsI8VxbGNNHGwJj2y2PJwArDNOF-vCz6tD6zgtZ-GidM58wXDnJfv3qwsehy78I4DkWURhEbrXg9BulxFJr9mOBlKjogCpEV199m1oGdcdfv6IzG5gNstwj2R1M43AC68IDKMiDRD54TB3BjC-6Y2ycpKU9QHsZl8fzhihIUG7ZOQ01Gt0fhO1JFyklR7ef1JVqp9yrdpqxfdQJNJwAlN3LzkbMnQyu7r-YPpItaiIMlgMOCK_yiWSs9-Zo8OPglN9yOow2NXC5NaOcUqKYNb2V9-gGqvggwQiSdpcNQED7yf8gwYUlFmFFmkhnvU09ZFB0wWZjq5dQIJROuXDvzTquwtlybkem2eiw4cIhE_5cCI2VY4NxeX7_Xhc9bXQ2qBYHe7zafWdiLdh_p-QcOMk0EmUtMpnfiL01aGcQYds9C3hayO51HxOKLA1ubDhuI0Kiw5M0sJ5QmZ1AeF99vGcf0HXOcnZVUifdh7R2xbXOjojGGz2T-ON1PbtnIRIUPDUYXvhwyVzE6VnFYYWnsnITIef32bfcUlgLFuARHrNDuPkQHR-GB2-aqfplFLJGc5rfpKTFgfISfUl8o-P0HNtP5VF8KYsu2lYkTWDcD0XNV6Ws_nDdPf8pR4otM5XvRKjkQC3xN3ks-5agFX63cjH8K4Ct2oTYHSYVNJxkaY8cZgzSbvBgRePWzsUvVenLbaHqww3PMxmmf1grg.w1Ce0nx5voBdupjkt7sk5w"
}
```
