import Fetch from "../fetch"

export interface Like {
  id: number
  user_id: number
  tweet_id: number
  created_at: string
}

export const likeAPI = {
  like: async (tweetId: number) => {
    const response = await Fetch.post<{ data: Like }>("/likes", { tweet_id: tweetId })
    return response.data.data
  },

  unlike: async (tweetId: number) => {
    await Fetch.delete(`/likes/${tweetId}`)
  },

  getTweetLikes: async (tweetId: number) => {
    const response = await Fetch.get<{ data: Like[] }>(`/likes/tweet/${tweetId}`)
    return response.data.data
  },
}
