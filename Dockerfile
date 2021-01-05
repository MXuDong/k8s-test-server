FROM golang:1.15

# Author info：
# - [ I don't like use MAINTAINER to define the Author]
# - Author : A-Donga
#
# Verwion : build-0.0.1

# Copy file
COPY app app

# Run program
CMD ./app