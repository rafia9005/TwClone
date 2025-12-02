import { useState, useCallback } from "react";
import { likeAPI, type Like } from "@/lib/api";

export const useLike = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [likes, setLikes] = useState<Like[]>([]);

  const likeTweet = useCallback(async (tweetId: number) => {
    setLoading(true);
    setError(null);
    try {
      const result = await likeAPI.like(tweetId);
      return result;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to like tweet";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const unlikeTweet = useCallback(async (tweetId: number) => {
    setLoading(true);
    setError(null);
    try {
      await likeAPI.unlike(tweetId);
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to unlike tweet";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const getTweetLikes = useCallback(async (tweetId: number) => {
    setLoading(true);
    setError(null);
    try {
      const data = await likeAPI.getTweetLikes(tweetId);
      setLikes(data);
      return data;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to fetch likes";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    loading,
    error,
    likes,
    likeTweet,
    unlikeTweet,
    getTweetLikes,
  };
};
