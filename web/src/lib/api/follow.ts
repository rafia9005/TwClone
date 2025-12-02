import Fetch from "../fetch"

export interface Follow {
  id: number
  follower_id: number
  following_id: number
  created_at: string
}

export const followAPI = {
  follow: async (userId: number) => {
    const response = await Fetch.post<{ data: Follow }>("/follows", { following_id: userId })
    return response.data.data
  },

  unfollow: async (userId: number) => {
    await Fetch.delete(`/follows/${userId}`)
  },

  getFollowers: async (userId: number) => {
    const response = await Fetch.get<{ data: Follow[] }>(`/follows/followers/${userId}`)
    return response.data.data
  },

  getFollowing: async (userId: number) => {
    const response = await Fetch.get<{ data: Follow[] }>(`/follows/following/${userId}`)
    return response.data.data
  },
}
