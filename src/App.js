import {
  BrowserRouter, Routes, Route,
} from 'react-router-dom';import './App.css';
import Home from './ReactComponent/Home';
import { useEffect, useState } from 'react';
import Footer from './ReactComponent/Footer/Footer';

function App() {

    const [message,setMessage] = useState("paf")

    const fetchData = async () => {
      const response = await fetch('http://localhost:8080/api/data');
      const data = await response.json();
      setMessage(data.message)
  };

    useEffect(() => {
      fetchData();
    },[])

  return (
    <div>
      {/* <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home/>} />
        </Routes>
      </BrowserRouter> */}
      <Footer/>
    </div>
   
  
  );
}

export default App;
