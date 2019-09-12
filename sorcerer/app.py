#!/usr/bin/env python
import json
from flask import Flask #, request

app = Flask(__name__)

@app.route("/")
def default_response():
    default_event = {'event_type': 'default'}
    return "This is the default response!\n"


@app.route("/solution")
def get_solution():
    get_solution_event = {'event_type': 'get_solution'}
    solution = {
        'angle': '42degrees',
        'mass': '16stone'
    }
    return json.dumps(solution)
