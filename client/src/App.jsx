import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import axios from 'axios'

function App() {
  const fetchAll = async ()=>{
    try {
      const res = await axios.get("http://localhost:4000/users") 
      console.log(res.data)
    } catch(err) {
      console.error("Unable to fetch all users", err)
    }
  }

  const fetchById = ()=>{1
    console.log("Clicked 2")
  }

  const handleClick3 = ()=>{1
    console.log("Clicked 3")
  }

  return (
    <div>
      <h1>Hello, this is a client that is consuming Go APIs</h1> 
      <div className="button-container">
        <button onClick={fetchAll}>fetch all users</button> 
        <button onClick={fetchById}>fetch user by id</button> 
        <button onClick={handleClick3}>create user</button> 
      </div> 
    </div> 
  )
}

export default App
