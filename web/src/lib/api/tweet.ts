import Fetch from "../fetch"
import type { User } from "./auth"

export interface Tweet {
  id: number
  user_id: number
  user?: User
  content: string
  media_url?: string
  created_at: string
  updated_at: string
  likes_count?: number
  replies_count?: number
  retweets_count?: number
  bookmarks_count?: number
}

export interface CreateTweetRequest {
  content: string
  media_url?: string
}

export interface UpdateTweetRequest {
  content?: string
  media_url?: string
}

export const tweetAPI = {
  getAll: async () => {
    const response = await Fetch.get<{ data: Tweet[] }>("/tweets")
    return response.data.data
  },

  getById: async (id: number) => {
    const response = await Fetch.get<{ data: Tweet }>(`/tweets/${id}`)
    return response.data.data
  },

  create: async (data: CreateTweetRequest) => {
    const response = await Fetch.post<{ data: Tweet }>("/tweets", data)
    return response.data.data
  },

  update: async (id: number, data: UpdateTweetRequest) => {
    const response = await Fetch.put<{ data: Tweet }>(`/tweets/${id}`, data)
    return response.data.data
  },

  delete: async (id: number) => {
    await Fetch.delete(`/tweets/${id}`)
  },
}
