import {useState} from "react"
import axios from "axios"
export default function CreateUserForm(){

  const [postData, setPostData] = useState({
    name: "",
    email: ""
  })

  const submitCUF = async (e) => {
    e.preventDefault()

    if (!postData.name || !postData.email){
      alert("Name and Email are required!!!")
      return      
    }

    try {
      const res = await axios.get("http://localhost:4000/create/user", postData, {
        headers: {
          "Content-Type": "application/json"
        }
      })
      console.log("User created:", res.data)
    } catch(err) {
      console.error("Error Creating User", err)
    }
  }

  const handleChange = (e)=>{
    const {name, value} = e.target;
    setPostData({
      ...postData,
      [name]: value
    })
  }

  return (
    <form onSubmit={submitCUF} className="cuf">
      <input onChange={handleChange} type="text" id="name" name="name" value={postData.name}/> 
      <input onChange={handleChange} type="email" id="email" name="email" value={postData.email}/> 
      <button type="submit">Create</button> 
    </form> 
  )
}
