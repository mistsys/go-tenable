#!/usr/bin/env python3
import os
import argparse

def parse_arguments():
    parser = argparse.ArgumentParser(description='go-tenable first pass integration tester')
    parser.add_argument("--accesskey", required=True, help="Access Key for Tenable.iO")
    parser.add_argument("--secretkey", required=True, help="Secret Key for Tenable.io")
    return parser.parse_args()

# TODO: replace with subprocess if we want to go this route
# or rewrite in Go
def basic_end_to_end(accesskey, secretkey):
    os.system("go install github.com/mistsys/go-tenable")
    os.system("go-tenable folders --accesskey {} --secretkey {} list".format(accesskey, secretkey))
    os.system("go-tenable scans --accesskey {} --secretkey {} list".format(accesskey, secretkey))
    os.system("go-tenable server --accesskey {} --secretkey {} status".format(accesskey,secretkey))

def main():
    args = parse_arguments()
    basic_end_to_end(args.accesskey, args.secretkey)

if __name__ == '__main__':
    main()
