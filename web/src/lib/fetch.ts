import axios from "axios"
import Cookies from "js-cookie"

const Fetch = axios.create({
  baseURL: `${import.meta.env.VITE_API}/api/v1`,
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
})

// Attach latest accessToken from cookie on every request
Fetch.interceptors.request.use((config) => {
  const token = Cookies.get("accessToken")
  if (token) {
    config.headers = config.headers || {}
    config.headers["Authorization"] = `Bearer ${token}`
  }
  return config
})

export default Fetch