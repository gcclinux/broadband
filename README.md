# Broadband (go)
----------------------------------------------------------------
#### The go Broadband app / code was writen to satisfy my own requirements to monitor my internet speed regularly

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