import grpc
import uvicorn
from fastapi import Depends, FastAPI, File, UploadFile
from proto.proto_pb2_grpc import ImageServiceStub
from proto.proto_pb2 import ConsumeRequest, ConsumeResponse, Image, ProduceRequest, ProduceResponse
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, update
from database import get_db, Thread
from compare import difference

app = FastAPI(debug=True)

client = ImageServiceStub(channel = grpc.insecure_channel('proc:8001', options=(('grpc.enable_http_proxy', 0),)))

def produce(image : list) -> ProduceResponse:
    return client.Produce(ProduceRequest(img=Image(value=image)))

def consume(img1 : str, img2 : str) -> ConsumeResponse:
    return client.Consume(ConsumeRequest(image1=img1, image2=img2))
    

@app.get("/health")
async def health():
    return {"status": "ok"}

@app.post("/create_thread")
async def create_thread(image: UploadFile = File(), db : AsyncSession = Depends(get_db)):
    image_name : ProduceResponse = produce(list(await image.read()))
    thread = Thread(image=image_name.image)
    db.add(thread)
    await db.flush()
    thread.thread = thread.id
    db.merge(thread)
    await db.commit()
    return {"thread" : thread.thread}

@app.post("/add_image/{thread}")
async def add_image(thread : int, image: UploadFile = File(), db : AsyncSession = Depends(get_db)):
    image_name : ProduceResponse = produce(list(await image.read()))
    db.add(Thread(image=image_name.image, thread=thread))
    await db.commit()
    return {"thread" : image_name.image}

@app.get("/get_thread/{thread}")
async def get_thread(thread : int, db : AsyncSession = Depends(get_db)):
    image = await db.execute(select(Thread).where(Thread.thread == thread).order_by(Thread.id.desc()).limit(10))
    image = image.scalars()
    if not image:
        return {"message" : "Thread not found"}
    image_list = []
    for x in image:
        image_list.append(f"http://localhost:9000/images/{x.image}")
    return {"images" : image_list}

@app.get("/diff/{image1}/{image2}")
async def diff(image1 : str, image2 : str):
    response : ConsumeResponse = consume(image1, image2)
        
    img = difference(bytes(response.image1.value), bytes(response.image2.value))
    
    url = produce(list(img))
    
    return {"diff" : url.image}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)