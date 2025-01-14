import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import axios from 'axios'
import CreateUserForm from "./Forms/CreateUserForm"

function App() {
  const [users, setUsers] = useState([])
  const [user, setUser] = useState(null)
  const [showCreateUserForm, setShowCreateUserForm] = useState(false)
  const fetchAll = async ()=>{
    try {
      const res = await axios.get("http://localhost:4000/users") 
      setUsers(res.data)
    } catch(err) {
      console.error("Unable to fetch all users", err)
    }
  }

  const fetchById = async ()=>{
    try {
      console.log("Fetching user with user id 1")
      const res = await axios.get("http://localhost:4000/user/1") 
      setUser(res.data)
    } catch(err) {
      console.error("Unable to fetch all users", err)
    }
  }

  const makeFormAppear = ()=>{
    setShowCreateUserForm(true)
  }

  return (
    <div>
      <h1>Hello, this is a client that is consuming Go APIs</h1> 
      <div className="button-container">
        <button onClick={fetchAll}>fetch all users</button> 
        <button onClick={fetchById}>fetch user by id</button> 
        <button onClick={makeFormAppear}>create user</button> 
      </div> 
      {/*fetching all users*/}
      <div>
        {users && users.map(user=>{
          return(
            <div key={user.id}>
              <p>{user.name}</p> 
              <p>{user.email}</p> 
            </div> 
          )
        })}
      </div> 
      {/*fetching user by id*/}
      <div>
        {user && (
          <div>
            <p>{user.name}</p> 
            <p>{user.email}</p> 
          </div> 
        )} 
      </div> 
      <div>
        {showCreateUserForm && <CreateUserForm/>}
      </div> 
    </div> 
  )
}

export default App
