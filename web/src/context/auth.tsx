import { createContext, useContext, useEffect, useState, type ReactNode } from "react"
import Cookies from "js-cookie"
import { userAPI, type User } from "@/lib/api"

type AuthContextProps = {
  user: User | null
  loading: boolean
  error: string | null
  refresh: () => Promise<void>
  updateUser: (user: User) => void
  logout: () => void
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined)

export function useAuth() {
  const ctx = useContext(AuthContext)
  if (!ctx) throw new Error("useAuth must be used within AuthProvider")
  return ctx
}

export function AuthProvider({ children, initialUser }: { children: ReactNode; initialUser?: User | null }) {
  const [user, setUser] = useState<User | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    if (initialUser) {
      setUser(initialUser)
      setLoading(false)
      refresh()
      return
    }
    refresh()
  }, [])

  const refresh = async () => {
    setLoading(true)
    setError(null)
    try {
      const accessToken = Cookies.get("accessToken")
      if (!accessToken) {
        setUser(null)
        setError("Unauthorized")
        setLoading(false)
        return
      }
      const userData = await userAPI.getMe()
      setUser(userData)
    } catch (err: any) {
      setUser(null)
      if (err.response && err.response.status === 401) {
        setError("Unauthorized")
      } else {
        setError("Network error")
      }
    } finally {
      setLoading(false)
    }
  }

  const updateUser = (newUser: User) => {
    setUser(newUser)
  }

  const logout = () => {
    Cookies.remove("accessToken")
    setUser(null)
  }

  return (
    <AuthContext.Provider value={{ user, loading, error, refresh, updateUser, logout }}>
      {children}
    </AuthContext.Provider>
  )
}