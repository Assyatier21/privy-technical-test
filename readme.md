# API Endpoint For Cakes

Cakes API provide user to create, read, update and delete cakes from the database.

| List Endpoints                                                            | Description                                        |
| ------------------------------------------------------------------------- | -------------------------------------------------- |
| [List of Cakes](https://www.notion.so/99bdf526c51c45fb8ae3384008dd97a5)   | Get List of Cakes Using Limit and Offset Parameter |
| [Details of Cake](https://www.notion.so/d8d6d469b0bd4f72a6361c58ebe75420) | Get Detail of Cakes By ID Param                    |
| [Add New Cake](https://www.notion.so/bb965f30aa1e4637b7892a3936717b5e)    | Add Cake Via Body Request                          |
| [Update Cake](https://www.notion.so/66003d12436a4cb180e35b1331895797)     | Update Cake Via Body Request                       |
| [Delete Cake](https://www.notion.so/1008980a065b42e0a9b7be686f0849ce)     | Delete Cake By ID Param                            |

## Installing and Running

### Locally:
```bash
$ go mod tidy
$ go run cmd/main.go
```

### Using Docker:

```bash
$ docker build --tag privy-technical-test .
$ docker run --rm -p 8800:8800 privy-technical-test
```
