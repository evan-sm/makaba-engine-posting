package main

import (
    "os"
)

const (
    makabaUrl  = "https://2ch.hk/makaba/makaba.fcgi"
    postingUrl = "https://2ch.hk/makaba/posting.fcgi?json=1"
)

var (
    passcode = os.Getenv("PASSCODE") // https://2ch.hk/2ch/
)
