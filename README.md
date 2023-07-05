# Govertime

Booky's unofficial overtime tracker

## Compile

```bash
go build
```

## Commands

```bash
# Record Overtime
./govertime [users] [start-date] [start-time] [end-date] [end-time] [reason]

# List Overtimes for the Month
./govertime -ls

# List Overtimes for the Month + Prev Month 29/30/31
./govertime -ls -gdtb

# List Overtimes for the Month + Prev Month 29/30/31 for a specific user
./govertime -ls -gdtb -names [users]

# Export to CSV
./govertime -e [path]

```
### Examples

```bash
./govertime "alec,wayne,luis" 07-03-2023 10:00PM 07-04-2023 12:00AM "Fixed bugs for GCash MP"
./govertime -ls -gdtb -names "cheks"
./govertime -ls -gdtb -names "abbae, geps, charles" -e "reports/june/DOS.csv"
```

### Params

|Param|Type|Required|Description|
|-----|----|--------|-----------|
|`users`|`string`|required|`"jett,abbae,geps"`, `"wayne"`|
|`[start\|end]-date`|`date`|required| `MM-DD-YYYY` eg. `01-23-2023`|
|`[start\|end]-time`|`time`|required| `hh:mmA` where `A` can be `AM` or `PM`|
|`reason`|`string`|required|`"DOS regression and bug fixes deployment"`|
|`path`|`string`|optional|directory for location csv. uses current directory by default|
