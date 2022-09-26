# Mini Incident CLI

Mini incident is a CLI tool written in the Go programming language that consumes incidents formatted in JSON and outputs it in CSV format,  in specified sort order, so it can be adapted in a web service output for a bulk DB insert.


## Installation

Download the most recent version of [go](https://go.dev/dl/)

clone this repo

```bash
git clone https://github.com/Vesino/GoCLIIncidents.git
```

## Usage

CLI arguments
```go
  -columns string
    	Columns to output in CSV
  -json-input string
    	JSON payload which contains incident data
  -sortdirection string
    	Sort columns in the specified direction, optional values: ascending or descending (default "ascending")
  -sortfield string
    	Sort columns by field, could, optional values: discovered or status (default "discovered")
```

```shell

BODY='[
{"id": 1,"name": "Misdirected fax","discovered": "2018-04-02","description": "Patients medical records faxed to wrong number.","status": "New"},
{"id": 4,"name": "Lost laptop","discovered": "2018-02-19","description": "Doctor lost her laptop while on vacation.","status": "Done"},
{"id": 2,"name": "Misdirected phone","discovered": "2018-04-02","description": "Patients medical records faxed to wrong number.","status": "New"},
{"id": 5,"name": "Lost iPad with medical record","discovered": "2018-04-01","description": "Nurse misplaced a patients medical record while in the office.","status": "In progress"},
{"id": 3,"name": "Misdirected email","discovered": "2018-04-02","description": "Patient´s medical records faxed to wrong number.","status": "New"}
]'


$ go run cmd/main.go -json-input=$BODY

$ go run cmd/main.go -json-input=$BODY -sortdirection='descending' -sortfield='status'

# Select specific columns
$ go run cmd/main.go -json-input=$BODY -sortdirection='descending' -sortfield='status' -columns='id, name, status'
```

## Tests
Execute test cases written 
```shell
go test ./... -v
```

## TODO
- Create and improve more Test Cases
- Be able to store the CSV in a given path

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
