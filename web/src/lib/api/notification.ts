import Fetch from "../fetch"

export interface Notification {
  id: number
  user_id: number
  type: string
  content: string
  is_read: boolean
  created_at: string
}

export const notificationAPI = {
  getAll: async () => {
    const response = await Fetch.get<{ data: Notification[] }>("/notifications")
    return response.data.data
  },

  markAsRead: async (id: number) => {
    await Fetch.put(`/notifications/${id}`, { is_read: true })
  },

  markAllAsRead: async () => {
    await Fetch.put("/notifications/read-all")
  },
}
