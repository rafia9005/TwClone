import { useState, useCallback } from "react";
import { authAPI, userAPI, type User } from "@/lib/api";
import Cookies from "js-cookie";

export const useAuth = () => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const login = useCallback(async (email: string, password: string) => {
    setLoading(true);
    setError(null);
    try {
      const response = await authAPI.login({ email, password });
      Cookies.set("accessToken", response.token, { expires: 7 });
      setUser(response.user);
      return response;
    } catch (err: any) {
      const message = err.response?.data?.message || "Login failed";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const register = useCallback(async (data: { name: string; username: string; email: string; password: string }) => {
    setLoading(true);
    setError(null);
    try {
      const user = await authAPI.register(data);
      setUser(user);
      return user;
    } catch (err: any) {
      const message = err.response?.data?.message || "Registration failed";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const logout = useCallback(() => {
    Cookies.remove("accessToken");
    setUser(null);
  }, []);

  const getMe = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const userData = await userAPI.getMe();
      setUser(userData);
      return userData;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to fetch user";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    user,
    loading,
    error,
    login,
    register,
    logout,
    getMe,
    setUser,
  };
};
