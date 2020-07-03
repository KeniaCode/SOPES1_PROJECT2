from flask import Flask, request, jsonify
from flask_restful import Resource, Api
import requests

app = Flask(__name__)
api = Api(app)


@app.route('/')
def hello_world():
    return 'Servidor inicializado'


class postData(Resource):
    def post(self):
        print("Recibiendo datos: ")
        print(request.json)
        # hacemosel POST a la base de datos
        response = requests.post('http://104.197.133.195:3000/addToredis', data=request.json)
        print("Respuesta de la base de datos: \n")
        print(response.text)
        return {'status': 'Nuevo dato agregado'}


api.add_resource(postData, '/postData')

if __name__ == '__main__':
     app.run(port='5000')