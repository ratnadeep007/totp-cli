# TOTP CLI

A simple CLI to get your TOTP, a cli replacement for Google Authenticator

#### Usage
- Download perferred version for your OS
- Move it to path eg `/usr/local/bin` and rename it to `totp`
- Add a file name `.totp` at your home directory
- Add content in format `<PROVIDER>=<2FA Secret Key>`
- You can add as many as you want
```
<PROVIDER_1>=<2FA Secret Key>
<PROVIDER_2>=<<2FA Secret Key>
```
- After adding run `totp`
- Output will like:
```
<PROVIDER_1>: <TOTP>
<PROVIDER_2>: <TOTP>
```