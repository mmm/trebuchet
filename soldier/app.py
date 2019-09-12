#!/usr/bin/env python
import requests
r = requests.get('http://sorcerer:5000/solution?distance=60m&wind=-10mps')
print(r.text)
