# Broadband (go)
----------------------------------------------------------------
#### The go Broadband app / code was writen to satisfy my own requirements to monitor my internet speed regularly

You can run this app/code as is and it will return the output on the screen or you have an option to store the data into a PostgreSQL database, with minor changes you can use other DBs like MySQL & SQLite but I have only included a config and tested on PostgreSQL

Output Example:

```
WanIP: 123.123.123.123, Latency: 0.013298, Download: 22.330900, Upload: 6.632200, (Manchester (United Kingdom) by Vodafone UK)
```

PostgreSQL Table:
```
CREATE TABLE IF NOT EXISTS public.broadband
(
    id SERIAL PRIMARY KEY,
    created timestamp without time zone NOT NULL DEFAULT now(),
    address character varying(32),  
    latency DOUBLE PRECISION,
    downloadspeed DOUBLE PRECISION,
    uploadspeed DOUBLE PRECISION,
    isp character varying(256)
)
```
To be able to store the data into your database you will need to create the above tables, add the details in the config and set SaveToDB to true!

Config Example:
```
{
    "DB_HOST":      [""],
    "DB_PORT":      [""],
    "DB_USER":      [""],
    "DB_PASS":      [""],
    "DB_NAME":      [""],
    "TB_NAME":      [""],
    "Verbose":      [true],
    "SaveToDB":     [false]
}
```
You can run by downloading the source code and execute the ```$./run.sh``` script you can also run it like this ```$ go run main.go``` or you can compile the source code into a binary like this ``` $ go build -o broadband-$(uname)-$(uname -m) *.go``` 

In the future I may add to run this as a service, add a graphic representation of the collected data, but this is just an initial idea for my current requirements.