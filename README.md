### Notes:
This did take longer than 30 minutes, probably more like 60-90. I understand that good incomplete answers were
acceptable, but I wanted to try and make it work.

I have a lot to learn and wouldn't be able to re-write something like this on the spot in a live whiteboarding environment

Thanks for coming to my TED talk

### Running the program
In the root folder, run `go run .`
The server should run and stay open on port `:8080`

I've only added one test to save time. This took longer than 30 minutes but, I can't in good conscious leave
my functions untested, it hurts my soul

run `go test -v ./...` to test

Output should look like:
```
=== RUN   TestGetMetersForCustomer
=== RUN   TestGetMetersForCustomer/Test_Case_1
=== RUN   TestGetMetersForCustomer/Test_Case_2
--- PASS: TestGetMetersForCustomer (0.00s)
    --- PASS: TestGetMetersForCustomer/Test_Case_1 (0.00s)
    --- PASS: TestGetMetersForCustomer/Test_Case_2 (0.00s)
PASS
```

I've tested the endpoints using the Insomnia REST client with the following GET Requests:

```localhost:8080/getMeterReading?serialID=1111-1111-2222```

result:
```
{
	"kWh reading": 11571
}
```

GET: `localhost:8080/getMetersForCustomer?customer=Aquaflow`

result:
```
[
	{
		"name": "Treatment Plant A",
		"customer": "Aquaflow",
		"serialID": "1111-1111-1111"
	},
	{
		"name": "Treatment Plant B",
		"customer": "Aquaflow",
		"serialID": "1111-1111-2222"
	}
]
```