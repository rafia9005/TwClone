import { useState, useCallback } from "react";
import { tweetAPI, type Tweet } from "@/lib/api";

export const useTweets = () => {
  const [tweets, setTweets] = useState<Tweet[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchTweets = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await tweetAPI.getAll();
      setTweets(data);
      return data;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to fetch tweets";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const getTweet = useCallback(async (id: number) => {
    setLoading(true);
    setError(null);
    try {
      const data = await tweetAPI.getById(id);
      return data;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to fetch tweet";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const createTweet = useCallback(async (content: string, mediaUrl?: string) => {
    setLoading(true);
    setError(null);
    try {
      const newTweet = await tweetAPI.create({ content, media_url: mediaUrl });
      setTweets((prev) => [newTweet, ...prev]);
      return newTweet;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to create tweet";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const updateTweet = useCallback(async (id: number, content: string) => {
    setLoading(true);
    setError(null);
    try {
      const updatedTweet = await tweetAPI.update(id, { content });
      setTweets((prev) =>
        prev.map((tweet) => (tweet.id === id ? updatedTweet : tweet))
      );
      return updatedTweet;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to update tweet";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const deleteTweet = useCallback(async (id: number) => {
    setLoading(true);
    setError(null);
    try {
      await tweetAPI.delete(id);
      setTweets((prev) => prev.filter((tweet) => tweet.id !== id));
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to delete tweet";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    tweets,
    loading,
    error,
    fetchTweets,
    getTweet,
    createTweet,
    updateTweet,
    deleteTweet,
    setTweets,
  };
};
