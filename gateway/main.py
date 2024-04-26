import grpc
import uvicorn
from fastapi import FastAPI, File, UploadFile
from proto.proto_pb2_grpc import ImageServiceStub
from proto.proto_pb2 import Image, ProduceRequest

app = FastAPI()

client = ImageServiceStub(channel=grpc.insecure_channel("localhost:8001"))

def produce(image : list):
    client.Produce(ProduceRequest(img=Image(value=image)))

@app.get("/health")
async def health():
    return {"status": "ok"}

@app.post("/add_image")
async def add_image(image: UploadFile = File()):
    produce(list(await image.read()))
    

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)