import argparse
import sys
import random

from numpy import byte

parser = argparse.ArgumentParser(description="Generate .lsrom files from binary input.")
parser.add_argument("-o", "--output", default=None, help="output file (default: out.lsrom)")
parser.add_argument("-I", "--stdin", action="store_true", help="read input from stdin")
parser.add_argument("-O", "--stdout", action="store_true", help="write output to stdout")
parser.add_argument("-r", "--rand", action="store_true", help="fill output with bytes random data")
parser.add_argument("-l", "--output-len", type=int, default=16777216, help="maximum length of output (default: 16777216)")
parser.add_argument("inputs", nargs="*", help='input files')

args = parser.parse_args()

input_bytes = []

output = None
if args.output != None:
    if args.stdout:
        print("-O / --stdout cannot be used when an output file is specified.");
        exit(1)

    output = open(args.output, "w+")
elif args.stdout:
    output = sys.stdout
else:
    output = open("out.lsrom", "w+")

if len(args.inputs) > 0:
    if args.stdin:
        print("-I / --stdin cannot be used when input files are specified.");
        exit(1)
    elif args.rand:
        print("-r / --rand cannot be used when input files are specified.");
        exit(1)

    else:
        for file in args.inputs:
            with open(file, "rb") as f:
                bytes_read = f.read()
                input_bytes.append(bytes_read)
                if len(input_bytes) > args.output_len:
                    print("Input files exceeded maximum output length")
                    exit(1)
elif args.stdin:
    if args.rand:
        print("-r / --rand cannot be used when reading from stdin.");
        exit(1)

    input_bytes = sys.stdin.read(args.output_len)

elif args.rand:
    input_bytes = random.randbytes(args.output_len)

else:
    print("No input files specified.")
    exit(1);

output_bytes = "v3.0 hex words plain\n"
byte_number = 0
for b in input_bytes:
    output_bytes += "{:02x}".format(b)

    byte_number += 1
    if byte_number == 16:
        byte_number = 0;
        output_bytes += "\n"
    else:
        output_bytes += " "

output.write(output_bytes)
output.close()