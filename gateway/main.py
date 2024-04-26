import uvicorn
from fastapi import FastAPI, File, UploadFile


app = FastAPI()

@app.get("/health")
async def health():
    return {"status": "ok"}

@app.post("/add_image")
async def add_image(image: UploadFile = File()):
    print(bytes(await image.read()))

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)