import grpc
import uvicorn
from fastapi import Depends, FastAPI, File, UploadFile
from proto.proto_pb2_grpc import ImageServiceStub
from proto.proto_pb2 import Image, ProduceRequest, ProduceResponse
from sqlalchemy.ext.asyncio import AsyncSession
from database import get_db, Thread

app = FastAPI()

client = ImageServiceStub(channel=grpc.insecure_channel("proc:8001"))

def produce(image : list):
    return client.Produce(ProduceRequest(img=Image(value=image)))

@app.get("/health")
async def health():
    return {"status": "ok"}

@app.post("/add_image")
async def add_image(image: UploadFile = File(), db : AsyncSession = Depends(get_db)):
    image_name : ProduceResponse = produce(list(await image.read()))
    print(image_name)
    db.add(Thread(image=image_name.image))
    db.commit()
    
@app.post("/check_image")
async def check_image():
    produce([1,2,3,4,5,6,7,8,9,10])

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)