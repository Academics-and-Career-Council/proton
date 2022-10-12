import os

from flask import Flask 
from flask import request, abort, jsonify, send_from_directory
import flask

from get_cpi_ap.main import get_cpi_ap

# UPLOAD_DIRECTORY = "/api_uploaded_files"

# if not os.path.exists(UPLOAD_DIRECTORY):
#    os.makedirs(UPLOAD_DIRECTORY)

app = Flask(__name__)


	
@app.route('/uploader', methods = ['GET', 'POST'])
def upload_file():
   if request.method == 'POST':
      f = request.files['file']
      # f.save(f.filename)
      grade_structure = get_cpi_ap(f)
      response = flask.jsonify(grade_structure)
      response.headers.add('Access-Control-Allow-Origin', '*')
      return response
    

if __name__ == "__main__":
    app.run(debug=True)