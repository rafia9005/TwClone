import { useState, useCallback } from "react";
import { followAPI, type Follow } from "@/lib/api";

export const useFollow = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [followers, setFollowers] = useState<Follow[]>([]);
  const [following, setFollowing] = useState<Follow[]>([]);

  const followUser = useCallback(async (userId: number) => {
    setLoading(true);
    setError(null);
    try {
      const result = await followAPI.follow(userId);
      return result;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to follow user";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const unfollowUser = useCallback(async (userId: number) => {
    setLoading(true);
    setError(null);
    try {
      await followAPI.unfollow(userId);
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to unfollow user";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const getFollowers = useCallback(async (userId: number) => {
    setLoading(true);
    setError(null);
    try {
      const data = await followAPI.getFollowers(userId);
      setFollowers(data);
      return data;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to fetch followers";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const getFollowing = useCallback(async (userId: number) => {
    setLoading(true);
    setError(null);
    try {
      const data = await followAPI.getFollowing(userId);
      setFollowing(data);
      return data;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to fetch following";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    loading,
    error,
    followers,
    following,
    followUser,
    unfollowUser,
    getFollowers,
    getFollowing,
  };
};
