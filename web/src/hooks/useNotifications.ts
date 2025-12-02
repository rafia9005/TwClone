import { useState, useCallback } from "react";
import { notificationAPI, type Notification } from "@/lib/api";

export const useNotifications = () => {
  const [notifications, setNotifications] = useState<Notification[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [unreadCount, setUnreadCount] = useState(0);

  const fetchNotifications = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await notificationAPI.getAll();
      setNotifications(data);
      setUnreadCount(data.filter((n) => !n.is_read).length);
      return data;
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to fetch notifications";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const markAsRead = useCallback(async (id: number) => {
    setLoading(true);
    setError(null);
    try {
      await notificationAPI.markAsRead(id);
      setNotifications((prev) =>
        prev.map((notif) =>
          notif.id === id ? { ...notif, is_read: true } : notif
        )
      );
      setUnreadCount((prev) => Math.max(0, prev - 1));
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to mark as read";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  const markAllAsRead = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      await notificationAPI.markAllAsRead();
      setNotifications((prev) =>
        prev.map((notif) => ({ ...notif, is_read: true }))
      );
      setUnreadCount(0);
    } catch (err: any) {
      const message = err.response?.data?.message || "Failed to mark all as read";
      setError(message);
      throw new Error(message);
    } finally {
      setLoading(false);
    }
  }, []);

  return {
    notifications,
    loading,
    error,
    unreadCount,
    fetchNotifications,
    markAsRead,
    markAllAsRead,
  };
};
