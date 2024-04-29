import { useState } from 'react';
import './App.css';
import {axiosInstance, MINIO} from './axios/axios';

function App() {
  const [img, setImg] = useState();
  const [thread, setThread] = useState(0);
  const [imgs, setImgs] = useState([]);
  const [diff, setDiff] = useState("");

  

  const createThread = async ()  => {
    const data = new FormData();
    console.log(img)
    data.append("image", img)

    let res = await axiosInstance.post("/create_thread", data)
    console.log(res.data);

    setImg(null)
    setThread(res.data.thread)
    console.log(thread)
  }

  const addImg = async (x)  => {
    const data = new FormData();
    console.log(img)
    data.append("image", img)

    let res = await axiosInstance.post("/add_image/" + x, data)
    console.log(res.data);

    setImg(null)
    console.log(thread)

    await getThread(x)
  }

  const getThread = async (thread) => {
    let res = await axiosInstance.get("/get_thread/" + thread)
    setImgs(res.data.images)
  }

  const getDiff = async () => {
    let res = await axiosInstance.get("/diff/"+imgs[imgs.length-1]+"/"+imgs[imgs.length-2])
    setDiff(res.data.diff)
  }

  return (
    <div className="App">
      <div>
        <input type='file' onChange={(e) => setImg(e.target.files[0])} />
        <button onClick={createThread}>создать тред</button>
      </div>
      <hr />

      {thread != 0 && 
        <div onClick={() => getThread(thread)}>
          <div>{thread}</div>
          <input type='file' onChange={(e) => setImg(e.target.files[0])} />
          <button onClick={() => addImg(thread)}>добавить изображение</button>
        </div>
      }   

      <hr />

      {imgs && imgs.map((img, index) => <img key={index} src={MINIO+img} className="max-image" />)}

      {imgs.length && <button onClick={() => getDiff()}>получить различия</button>}
      
      {diff && <img src={MINIO + diff} className="max-image" />}

    </div>
  );
}

export default App;
