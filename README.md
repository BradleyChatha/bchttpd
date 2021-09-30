# Overview

This is an ultra lightweight HTTP server.

All it supports is HTTP GET, and that's by design because I'll be running this on memory constrained machines that simply
serve static files.

I used Go because I wanted a tiny footprint with easy and efficient networking, which Go excels at.

It uses a Radix tree for routing.