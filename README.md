# Password

## Installation

```
go get -u github.com/gouniverse/password
```

## Usage

- Hash a password

```
password.Hash(password)
```

- Verify a password against a hash
```
password.Verify(password, hash)
```
