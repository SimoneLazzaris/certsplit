# certsplit
Extract a certificate from a .pem file containing multiple certificates, allowing quick access to the n-th one.

## Why ?

If you deal with x509 certificates, you'll find yourself handling files in PEM having multiple certificate chained. 
And maybe you want to access the 7-th one in a list of 29.

This simple tool can help you: it extracts the desired ascii encoded PEM certificate from the file, so you can parse it via openssl or other tools.

## Usage:

Usage of certsplit:

    certsplit [-n] [filename]
      -debug
            enable debugging
      -n int
            certificate number (default 1)

If you don't specifiy the filename, it goes with stdin.
