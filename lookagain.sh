#! /bin/bash

find . -name '*.sh' | sed 's#/##g' | sed  's/test//g' | cut -d  '.' -f2 | cut -d 'h' -f2