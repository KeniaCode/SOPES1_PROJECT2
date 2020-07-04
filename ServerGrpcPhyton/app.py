from flask import Flask, request, jsonify
from flask_restful import Resource, Api
from concurrent import futures
import grpc
import time
import caso_pb2_grpc
import caso_pb2 as caso_messages


app = Flask(__name__)
api = Api(app)



class InsertarService(caso_pb2_grpc.CasoServicer):
    def obtenerDatos(self, request, context):
        #response = datos_pb2.Respuesta()
        print(request.json)
        return "Dato insertado"




if __name__ == '__main__':
    print('Servidor gRPC en puerto: 50051')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    caso_pb2_grpc.add_CasoServicer_to_server(InsertarService(), server)
    server.add_insecure_port('[::]:5000')
    server.start()
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)