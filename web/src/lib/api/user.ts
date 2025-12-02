import Fetch from "../fetch"
import type { User } from "./auth"

export const userAPI = {
  getAll: async () => {
    const response = await Fetch.get<{ data: User[] }>("/users")
    return response.data.data
  },

  getById: async (id: number) => {
    const response = await Fetch.get<{ data: User }>(`/users/${id}`)
    return response.data.data
  },

  getMe: async () => {
    const response = await Fetch.get<{ data: User }>("/users/token")
    return response.data.data
  },

  update: async (id: number, data: Partial<User>) => {
    const response = await Fetch.put<{ data: User }>(`/users/${id}`, data)
    return response.data.data
  },

  delete: async (id: number) => {
    await Fetch.delete(`/users/${id}`)
  },
}
