from server import app
from flask import render_template


@app.route('/handle')
def handle():
    return render_template('handle.html')
