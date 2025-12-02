// Export all API modules
export * from "./auth"
export * from "./user"
export * from "./tweet"
export * from "./follow"
export * from "./like"
export * from "./notification"
export * from "./media"

// Re-export for convenient access
import { authAPI } from "./auth"
import { userAPI } from "./user"
import { tweetAPI } from "./tweet"
import { followAPI } from "./follow"
import { likeAPI } from "./like"
import { notificationAPI } from "./notification"
import { mediaAPI } from "./media"

export const api = {
  auth: authAPI,
  user: userAPI,
  tweet: tweetAPI,
  follow: followAPI,
  like: likeAPI,
  notification: notificationAPI,
  media: mediaAPI,
}

export default api
