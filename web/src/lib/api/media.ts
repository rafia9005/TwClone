import Fetch from "../fetch"

export interface Media {
  id: number
  tweet_id: number
  media_url: string
  media_type: string
  created_at: string
}

export const mediaAPI = {
  upload: async (file: File) => {
    const formData = new FormData()
    formData.append("file", file)
    
    const response = await Fetch.post<{ data: Media }>("/media", formData, {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    })
    return response.data.data
  },

  getByTweetId: async (tweetId: number) => {
    const response = await Fetch.get<{ data: Media[] }>(`/media/tweet/${tweetId}`)
    return response.data.data
  },

  delete: async (id: number) => {
    await Fetch.delete(`/media/${id}`)
  },
}
