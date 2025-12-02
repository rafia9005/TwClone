import { useState, useCallback } from "react";
import { userAPI, type User } from "@/lib/api";

export const useUser = () => {
  const [user, setUser] = useState<User | null>(null);
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchUser = useCallback(async (id: number) => {
    setLoading(true);
    setError(null);
    try {
      const data = await userAPI.getById(id);
      setUser(data);
      return data;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to fetch user";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const fetchUsers = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await userAPI.getAll();
      setUsers(data);
      return data;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to fetch users";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const updateUser = useCallback(async (id: number, data: Partial<User>) => {
    setLoading(true);
    setError(null);
    try {
      const updatedUser = await userAPI.update(id, data);
      setUser(updatedUser);
      return updatedUser;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to update user";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const deleteUser = useCallback(async (id: number) => {
    setLoading(true);
    setError(null);
    try {
      await userAPI.delete(id);
      setUser(null);
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to delete user";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    user,
    users,
    loading,
    error,
    fetchUser,
    fetchUsers,
    updateUser,
    deleteUser,
    setUser,
  };
};
