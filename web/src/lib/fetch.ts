import axios from "axios"
import Cookies from "js-cookie"

const Fetch = axios.create({
  baseURL: `${import.meta.env.VITE_API}/api/v1`,
  headers: {
    "Content-Type": "application/json",
    Authorization: `Bearer ${Cookies.get("accessToken")}`
  },
})

export default Fetch