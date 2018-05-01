# EncDec

Encoder/decoder utilities for some of the more popular string
encoding utilities.

Supports the following schemes:
- Base64
- Base32
- URL
- Hex

## Quick Start

Run make.  Run 'enc -e <encoding> <text>' to encode the text.  Run 'dec -d <encoding> <text>' to decode the text.

If no text is specified on the command line, standard input is used.  This allows you to run piped encoding/decodings.

Example: enc -s b64 "Hello world!" | dec -s b64