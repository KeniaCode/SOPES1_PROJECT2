from concurrent import futures
import logging
from google.protobuf.json_format import MessageToDict
import grpc
import time
import caso_pb2_grpc
import caso_pb2


class CasoService(caso_pb2_grpc.CasoServicer):
    def CrearCasos(self, request, context):
        response = MessageToDict(request, preserving_proto_field_name=True)
        print("Recibido: ")
        print(response)
        respuesta = caso_pb2.CasoReply(mensaje='Datos recibidos')
        return respuesta


if __name__ == '__main__':
    print('Starting server. Listening on port 5000')
    logging.basicConfig()
    logging.info('Starting server. Listening')
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=20))
    caso_pb2_grpc.add_CasoServicer_to_server(CasoService(), server)
    server.add_insecure_port('[::]:5000')
    server.start()
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)
