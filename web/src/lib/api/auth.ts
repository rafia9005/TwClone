import Fetch from "../fetch"

export interface User {
  id: number
  email: string
  name: string
  username: string
  avatar?: string
  banner?: string
  bio?: string
  verified?: boolean
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  email?: string
  username?: string
  password: string
}

export interface RegisterRequest {
  email: string
  name?: string
  username: string
  avatar?: string
  banner?: string
  bio?: string
  password: string
}

export interface AuthResponse {
  token: string
  user: User
}

export const authAPI = {
  login: async (data: LoginRequest) => {
    const response = await Fetch.post<{ data: AuthResponse }>("/auth/login", data)
    return response.data.data
  },

  register: async (data: RegisterRequest) => {
    const response = await Fetch.post<{ data: User }>("/auth/register", data)
    return response.data.data
  },
}
