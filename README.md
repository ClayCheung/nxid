# NXID: Generate XID by nanosecond

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ClayCheung/nxid) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ClayCheung/nxid/master/LICENSE)

Package nxid is a nanosecond-level globally unique ID generator library, derived from [XID](https://github.com/rs/xid). NXID is specifically designed for scenarios requiring high-resolution timestamp precision.

## Key Changes from Original XID

- Utilizes an 8-byte nanosecond-level timestamp instead of a 4-byte second-level timestamp.
- Better suited for time-sensitive ID generation, ensuring higher resolution timing.

## Structure

- 8-byte value representing the nanoseconds since the Unix epoch,
- 3-byte machine identifier,
- 2-byte process ID, and
- 3-byte counter, starting with a random value.

The binary representation of the ID remains compatible with MongoDB's 12-byte Object IDs, but the time precision is significantly higher.

## Advantages

- Higher time resolution, making it suitable for applications where timestamp precision is essential.
- Compact and URL-safe string representation (20 chars).

## Comparison

| Name          | Binary Size | String Size    | Features
|---------------|-------------|----------------|----------------
| UUID          | 16 bytes    | 36 chars       | configuration-free, not sortable
| shortuuid     | 16 bytes    | 22 chars       | configuration-free, not sortable
| Snowflake     | 8 bytes     | up to 20 chars | requires machine/DC configuration, central server, sortable
| MongoID       | 12 bytes    | 24 chars       | configuration-free, sortable
| xid           | 12 bytes    | 20 chars       | configuration-free, sortable, second-level precision
| **nxid**      | 16 bytes    | 26 chars       | configuration-free, sortable, nanosecond-level precision

## Features

- Base32 hex encoded (20 chars when transported as a printable string, still sortable)
- Configuration-free, no need for unique machine or data center IDs
- K-ordered
- Embedded time with nanosecond precision
- Uniqueness guaranteed for 16,777,216 (24 bits) unique IDs per nanosecond and per host/process
- Lock-free generation

## Installation

```bash
go get github.com/ClayCheung/nxid
```

## Usage

Generate a new NXID:

```go
guid := nxid.New()

println(guid.String())
// Output example: 2vmide3burdh0kb22i1m0pln98
```

Retrieve embedded info:

```go
guid.Machine()
guid.Pid()
guid.Time()
guid.Counter()
```

## Licensing

All source code is licensed under the [MIT License](https://raw.github.com/ClayCheung/nxid/master/LICENSE).

---

### References

- Original XID: https://github.com/rs/xid