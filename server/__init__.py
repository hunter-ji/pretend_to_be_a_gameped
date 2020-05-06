from flask import Flask

app = Flask('server')
app.config.from_pyfile('config.py')

from server.views import home, handle